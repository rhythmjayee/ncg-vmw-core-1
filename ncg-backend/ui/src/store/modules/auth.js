import axios from 'axios';

const state = {
  user: {},
  functions: null,
};

const getters = {
  isAuthenticated: (state) => {
    console.log(state);
    return !!state.user;
  },
  StateFunctions: (state) => state.functions,
  StateUser: (state) => state.user,
};
const actions = {
  async Register({ dispatch }, form) {
    await axios.post('register', form);
    const UserForm = new FormData();
    UserForm.append('username', form.username);
    UserForm.append('email', form.email);
    UserForm.append('password', form.password);
    await dispatch('LogIn', UserForm);
  },
  async LogIn({ commit }, User) {
    await axios.post('login', User);
    const usr = {};
    usr.email = User.get('email');
    await commit('setUser', usr);
  },
  async GetFunctions({ commit }) {
    const response = await axios.get('functions');
    commit('setFunctions', response.data);
  },
  async LogOut({ commit }) {
    const user = null;
    commit('LogOut', user);
  },
};
const mutations = {
  setUser(state, user) {
    state.user = user;
  },
  setFunctions(state, functions) {
    state.functions = functions;
  },
  LogOut(state) {
    state.user = null;
    state.functions = null;
  },
};
export default {
  state,
  getters,
  actions,
  mutations,
};
