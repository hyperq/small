import Vue from 'vue';
import {
  InfiniteScroll,
  Spinner,
  Tabbar,
  TabItem,
  Field,
  Toast,
  Indicator,
  Actionsheet,
  Swipe,
  Switch,
  SwipeItem,
  Badge,
  CellSwipe,
  MessageBox,
  Checklist
} from 'mint-ui';
import 'mint-ui/lib/style.css';
Vue.prototype.$Toast = Toast;
Vue.prototype.$MessageBox = MessageBox;
Vue.prototype.$Loading = Indicator;
Vue.use(InfiniteScroll);
Vue.component(Spinner.name, Spinner);
Vue.component(Tabbar.name, Tabbar);
Vue.component(TabItem.name, TabItem);
Vue.component(Field.name, Field);
Vue.component(Actionsheet.name, Actionsheet);
Vue.component(Swipe.name, Swipe);
Vue.component(SwipeItem.name, SwipeItem);
Vue.component(Switch.name, Switch);
Vue.component(Badge.name, Badge);
Vue.component(CellSwipe.name, CellSwipe);
Vue.component(Checklist.name, Checklist);
