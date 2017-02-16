require('bootstrap/dist/css/bootstrap.min');
require('bootstrap/dist/js/bootstrap.min');
require('animate.css/animate.min.css');
require('../css/app');

import 'filters';
import TaskList from './components/task-list';
import SendTask from './components/send-task';
import TaskDetail from './components/task-detail';
import LoginModal from './components/login';

Vue.use(VueResource);
Vue.use(VueRouter);

const routes = [{
    path: '/',
    components: {
        default: TaskList,
        login: LoginModal,
    }
}, {
    path: '/tasks',
    components: {
        default: TaskList,
        login: LoginModal,
    }
}, {
    path: '/send',
    components: {
        default: SendTask,
        login: LoginModal,
    }
}, {
    path: '/detail/:id',
    components: {
        default: TaskDetail,
        login: LoginModal,
    }
}];

const router = new VueRouter({
    mode: 'history',
    routes: routes
});

const app = new Vue({
    router,
}).$mount('#main');
