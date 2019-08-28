import Vue from 'vue';
import Tarbar from '@/components/base/Tarbar';
import { areaData } from '@/utils/area';

Vue.component(Tarbar.name, Tarbar);

//微信分享
// import { wxRegister, wxShare, setNavigation } from '@/utils/wx';
Vue.prototype.$area = areaData();
