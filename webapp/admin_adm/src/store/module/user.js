import {
  getContentByMsgId,
  getMessage,
  getUnreadCount,
  getUserInfo,
  hasRead,
  login,
  logout,
  removeReaded,
  restoreTrash,
} from '@/api/user';

import {UserMutation} from '../mutation_constant';
import {getToken, setToken} from '@/libs/util';

const t = {
  namespaced: false,
};
t.state = {
  userName: '',
  userId: '',
  avatarImgPath: '',
  token: getToken(),
  access: '',
  hasGetInfo: false,
  unreadCount: 0,
  messageUnreadList: [],
  messageReadedList: [],
  messageTrashList: [],
  messageContentStore: {},
};
t.mutations = {
  [UserMutation.setAvatar](state, avatarPath) {
    state.avatarImgPath = avatarPath;
  },
  [UserMutation.setUserId](state, id) {
    state.userId = id;
  },
  [UserMutation.setUserName](state, name) {
    state.userName = name;
  },
  [UserMutation.setAccess](state, access) {
    state.access = access;
  },
  [UserMutation.setToken](state, token) {
    state.token = token;
    setToken(token);
  },
  [UserMutation.setHasGetInfo](state, status) {
    state.hasGetInfo = status;
  },
  [UserMutation.setMessageCount](state, count) {
    state.unreadCount = count;
  },
  [UserMutation.setMessageUnreadList](state, list) {
    state.messageUnreadList = list;
  },
  [UserMutation.setMessageReadedList](state, list) {
    state.messageReadedList = list;
  },
  [UserMutation.setMessageTrashList](state, list) {
    state.messageTrashList = list;
  },
  [UserMutation.updateMessageContentStore](state, {msg_id, content}) {
    state.messageContentStore[msg_id] = content;
  },
  [UserMutation.moveMsg](state, {from, to, msg_id}) {
    const index = state[from].findIndex(_ => _.msg_id === msg_id);
    const msgItem = state[from].splice(index, 1)[0];
    msgItem.loading = false;
    state[to].unshift(msgItem);
  },
};
t.getters = {
  messageUnreadCount: state => state.messageUnreadList.length,
  messageReadedCount: state => state.messageReadedList.length,
  messageTrashCount: state => state.messageTrashList.length,
};
t.actions = {};
t.actions.handleLogin = ({commit}, {userName, password}) => {
  userName = userName.trim();
  return new Promise((resolve, reject) => {
    login({
      userName,
      password,
    }).then(res => {
      const data = res.data;
      commit(UserMutation.setToken, data.token);
      resolve();
    }).catch(err => {
      reject(err);
    });
  });
};
t.actions.handleLogOut = ({state, commit}) => {
  return new Promise((resolve, reject) => {
    logout(state.token).then(() => {
      commit(UserMutation.setToken, '');
      commit(UserMutation.setAccess, []);
      resolve();
    }).catch(err => {
      reject(err);
    });
  });
};
t.actions.getUserInfo = ({state, commit}) => {
  return new Promise((resolve, reject) => {
    try {
      getUserInfo(state.token).then(res => {
        const data = res.data;
        commit(UserMutation.setAvatar, data.avatar);
        commit(UserMutation.setUserName, data.name);
        commit(UserMutation.setUserId, data.user_id);
        commit(UserMutation.setAccess, data.access);
        commit(UserMutation.setHasGetInfo, true);
        resolve(data);
      }).catch(err => {
        reject(err);
      });
    } catch (error) {
      reject(error);
    }
  });
};

t.actions.getUnreadMessageCount = ({state, commit}) => {
  getUnreadCount().then(res => {
    const {data} = res;
    commit(UserMutation.setMessageCount, data);
  });
};

t.actions.getMessageList = ({state, commit}) => {
  return new Promise((resolve, reject) => {
    getMessage().then(res => {
      const {unread, readed, trash} = res.data;
      commit(UserMutation.setMessageUnreadList, unread.sort(
        (a, b) => new Date(b.create_time) - new Date(a.create_time)));
      commit(UserMutation.setMessageReadedList, readed.map(_ => {
        _.loading = false;
        return _;
      }).sort((a, b) => new Date(b.create_time) - new Date(a.create_time)));
      commit(UserMutation.setMessageTrashList, trash.map(_ => {
        _.loading = false;
        return _;
      }).sort((a, b) => new Date(b.create_time) - new Date(a.create_time)));
      resolve();
    }).catch(error => {
      reject(error);
    });
  });
};
t.actions.getContentByMsgId = ({state, commit}, {msg_id}) => {
  return new Promise((resolve, reject) => {
    let contentItem = state.messageContentStore[msg_id];
    if (contentItem) {
      resolve(contentItem);
    } else {
      getContentByMsgId(msg_id).then(res => {
        const content = res.data;
        commit(UserMutation.updateMessageContentStore, {msg_id, content});
        resolve(content);
      });
    }
  });
};
t.actions.hasRead = ({state, commit}, {msg_id}) => {
  return new Promise((resolve, reject) => {
    hasRead(msg_id).then(() => {
      commit(UserMutation.moveMsg, {
        from: 'messageUnreadList',
        to: 'messageReadedList',
        msg_id,
      });
      commit(UserMutation.setMessageCount, state.unreadCount - 1);
      resolve();
    }).catch(error => {
      reject(error);
    });
  });
};

t.actions.removeReaded = ({commit}, {msg_id}) => {
  return new Promise((resolve, reject) => {
    removeReaded(msg_id).then(() => {
      commit('moveMsg', {
        from: 'messageReadedList',
        to: 'messageTrashList',
        msg_id,
      });
      resolve();
    }).catch(error => {
      reject(error);
    });
  });
};

t.actions.restoreTrash = ({commit}, {msg_id}) => {
  return new Promise((resolve, reject) => {
    restoreTrash(msg_id).then(() => {
      commit('moveMsg', {
        from: 'messageTrashList',
        to: 'messageReadedList',
        msg_id,
      });
      resolve();
    }).catch(error => {
      reject(error);
    });
  });
};

export default t;
