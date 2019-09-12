import store from '@/store';
import {getToken, setToken} from '@/libs/util';
import config from '@/config';

const LOGIN_PAGE_NAME = 'login';
const {homeName} = config;
var indexedRoutes = [];

const collectIndexedRoute = (routers, list, parent = '') => {
  routers.forEach(r => {
    if (r.children && r.children.length) {
      collectIndexedRoute(r.children, list, r.name);
    }
    list[r.name] = {
      access: r.meta && r.meta.access ? r.meta.access : '',
      parent: parent,
    };
  });

  return list;
};

const canAccess = (route, access) => {
  if (route.access) {
    return access.indexOf(route.access) !== -1;
  }

  return true;
};
const canTurnTo = (name, access, routes) => {
  const r = routes[name];
  if (r.parent && !canAccess(routes[r.parent], access)) {
    return false;
  }

  return canAccess(r, access);
};

const turnTo = (to, access) => {
  if (canTurnTo(to, access, indexedRoutes)) {
    return null;
  }

  return {replace: true, name: 'error_401'};
};

const accessPage = (to, next) => {
  if (store.state.user.hasGetInfo) {
    let turn = turnTo(to, store.state.user.access);
    if (turn) {
      next(turn);
    } else {
      next();
    }
    return;
  }
  store.dispatch('getUserInfo').then(user => {
    let turn = turnTo(to, store.state.user.access);
    if (turn) {
      next(turn);
    } else {
      next();
    }
  }).catch(() => {
    setToken('');
    next({name: LOGIN_PAGE_NAME});
  });
};

export const init = (routers) => {
  indexedRoutes = collectIndexedRoute(routers, []);
};

export const accessRoute = (to, from, next) => {
  const token = getToken();
  if (token) {
    if (to.name === LOGIN_PAGE_NAME) {
      next({name: homeName});
      return;
    }
    accessPage(to.name, next);
  } else if (to.name === LOGIN_PAGE_NAME) {
    next();
  } else {
    next({name: LOGIN_PAGE_NAME});
  }
};
