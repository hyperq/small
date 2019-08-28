import Vue from 'vue';

import '@/assets/css/public.css';
import '@/assets/icon/iconfont.css';
//ajax请求
import { get, post } from '@/utils/ajax';
Vue.prototype.$get = get;
Vue.prototype.$post = post;
//后台地址
Vue.prototype.$url = 'http://localhost';
