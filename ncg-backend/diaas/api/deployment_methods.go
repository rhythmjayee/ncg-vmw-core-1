package api

import (
	. "DIaaS/diaas/constants"
	. "DIaaS/diaas/db"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Deployment) Run(stdin io.Reader, stdout, stderr io.Writer, params string, duration time.Duration) error {
	if !d.Ready {
		return fmt.Errorf("failed to start: function not ready")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DeploymentTimeout)
	queryParams := strings.Split(params, " ")
	args := append([]string{"run", d.Image}, queryParams...)
	cmd := exec.CommandContext(ctx, "docker", args...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	time.Sleep(duration)
	err := cmd.Run()
	cancel()
	if err != nil {
		return fmt.Errorf("failed to run: %v", err)
	}
	return nil
}

func (d *Deployment) Build() error {
	var db *mongo.Database

	db = DbConnection()
	dir := Dirs + d.ID
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to allocate dir: %v", err)
	}

	err = ioutil.WriteFile(filepath.Join(dir, "main."+d.Language), d.Source, 0644)
	if err != nil {
		return fmt.Errorf("failed to write source: %v", err)
	}

	err = ioutil.WriteFile(filepath.Join(dir, "Dockerfile"), Dockerfiles[d.Language], 0644)
	if err != nil {
		return fmt.Errorf("failed to write dockerfile: %v", err)
	}

	imageName := fmt.Sprintf("dias-%16.16s", d.ID)
	log.Printf(imageName)
	cmd := exec.Command("docker", "build", "-t", imageName, dir)
	go func() {
		tmpfile, err := ioutil.TempFile("", "dias-*.log")
		if err != nil {
			log.Printf("Failed to hookup build logs: %v\n", err)
			return
		}
		d.BuildLogs = tmpfile.Name()
		cmd.Stdout = tmpfile
		cmd.Stderr = tmpfile
		err = cmd.Run()
		if err != nil {
			log.Printf("Failed to build %s (%s): %v\n", d.Image, d.Language, err)
		} else {
			coll := db.Collection("registry")
			log.Printf(d.Image)
			filter := bson.D{{"image", d.Image}}
			update := bson.D{{"$set", bson.D{{"ready", true}}}}
			log.Printf("Built image %s (%s)\n", d.Image, d.Language)
			result, err := coll.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				panic(err)
			}
			log.Printf("Number of documents updated: %v\n", result.ModifiedCount)
			d.Ready = true
		}
	}()
	return nil
}
