import Vue from 'vue';
import Router from 'vue-router';
import routes from './routers';
import iView from 'iview';
import {accessRoute, init} from '@/middleware/permission_route';
import {setTitle} from '@/libs/util';

Vue.use(Router);

const router = new Router({
  routes,
  mode: 'history',
});

init(routes);

router.beforeEach((to, from, next) => {
  iView.LoadingBar.start();
  console.log(to.name)
  accessRoute(to, from, next);
});

router.afterEach(to => {
  setTitle(to, router.app)
  iView.LoadingBar.finish()
  window.scrollTo(0, 0)
})

export default router;
