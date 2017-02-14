require('bootstrap/dist/css/bootstrap.min');
require('bootstrap/dist/js/bootstrap.min');
require('../css/app');

import 'filters';
import TaskList from './components/task-list';
import SendTask from './components/send-task';
import TaskDetail from './components/task-detail';

Vue.use(VueResource);
Vue.use(VueRouter);

const routes = [{
    path: '/tasks',
    component: TaskList,
}, {
    path: '/send',
    component: SendTask,
}, {
    path: '/detail/:id',
    component: TaskDetail,
}];

const router = new VueRouter({
    mode: 'history',
    routes: routes
});

const app = new Vue({
    router,
}).$mount('#main');
