import axios from 'axios';

const state = {
  user: null,
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
  async Register({ dispatch }, User) {
    await axios.post('register', User);
    await dispatch('LogIn', User);
  },
  async LogIn({ commit }, User) {
    await axios.post('login', User);
    await commit('setUser', User.username);
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
