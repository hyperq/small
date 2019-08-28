<template>
  <div class="tarbar">
    <div :class="index==1?'select':''"
         @click="change(1)">
      <div>
        <i class="iconfont icon-shouye"></i>
      </div>
      首页
    </div>
    <div :class="index==2?'select':''"
         @click="change(2)">
      <div>
        <i class="iconfont icon-fenleisousuo"></i>
      </div>
      分类
    </div>
    <div :class="index==3?'select':''"
         @click="change(3)">
      <div>
        <i class="iconfont icon-faxian"></i>
      </div>
      发现
    </div>
    <div :class="index==4?'select':''"
         @click="change(4)">
      <div>
        <i class="iconfont icon-gouwuche"></i>
      </div>
      购物车
    </div>
    <div :class="index==5?'select':''"
         @click="change(5)">
      <div>
        <i class="iconfont icon-wodejuhuasuan"></i>
      </div>
      我的
    </div>
  </div>
</template>

<script>
export default {
  name: "my-tarbar",
  data() {
    return {
      index: this.selected
    };
  },
  props: {
    selected: Number
  },
  methods: {
    change(index) {
      if (index != this.index) {
        let url = "";
        switch (parseInt(index)) {
          case 1:
            url = `/index`;
            break;
          case 2:
            url = `/column`;
            break;
          case 3:
            url = `/find`;
            break;
          case 4:
            url = `/shopcart`;
            break;
          case 5:
            url = `my`;
            break;
        }
        this.$router.push(url);
      }
    }
  }
};
</script>

<style scoped>
.tarbar {
  height: 55px;
  padding: 5px 0;
  position: fixed;
  bottom: 0;
  left: 0;
  background-color: #fff;
  width: 100%;
  overflow: hidden;
  z-index: 100;
  box-shadow: 0 0 10px 0 hsla(0, 6%, 58%, 0.6);
  font-size: 12px;
}
.tarbar > div div {
  line-height: 30px;
  height: 30px;
}
.tarbar > div {
  width: 20%;
  float: left;
  text-align: center;
  font-size: 12px;
  color: #666;
  position: relative;
}

.tarbar > div i {
  font-size: 24px;
}
.select,
.select div,
.select i {
  color: #e43130 !important;
}
</style>

