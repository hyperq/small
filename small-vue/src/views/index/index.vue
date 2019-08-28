<template>
  <div id="tab">
    <div class="floor slider-wrapper">
      <div class="header">
        <div class="header-icon">
          <i class="iconfont icon-fenleisousuo"></i>
        </div>
        <div class="header-input">
          <i class="iconfont icon-search"></i>
          <div>
            <input type="search"
                   placeholder="请输入要搜索的商品">
          </div>
        </div>
        <div class="header-icon">
          <i class="iconfont icon-more"></i>
        </div>
      </div>
      <div class="slider-bg"></div>
      <div class="new-slide">
        <cube-slide ref="slide"
                    :data="carousel"
                    @change="changePage">
          <cube-slide-item v-for="(item, index) in carousel"
                           :key="index"
                           @click.native="clickHandler(item, index)">
            <router-link :to="item.url">
              <img :src="item.img">
            </router-link>
          </cube-slide-item>
        </cube-slide>
      </div>
    </div>
    <div class="feature">
      <div></div>
      <div></div>
      <div></div>
      <div></div>
    </div>
    <my-tarbar :selected="1" />
  </div>
</template>

<script>
export default {
  data() {
    return {
      carousel: [
        {
          img:
            "http://m.360buyimg.com/mobilecms/s700x280_jfs/t1/38364/38/13953/98600/5d51359eE5caf4a1a/4a3dc7e660995817.jpg!cr_1125x549_0_72!q70.jpg.dpg",
          url: ""
        },
        {
          img:
            "http://m.360buyimg.com/mobilecms/s700x280_jfs/t1/10798/34/12141/189137/5c7668a1Eef6e08ec/1b7434cc17f6306d.jpg!cr_1125x549_0_72!q70.jpg.dpg",
          url: ""
        },
        {
          img:
            "http://m.360buyimg.com/mobilecms/s700x280_jfs/t1/10798/34/12141/189137/5c7668a1Eef6e08ec/1b7434cc17f6306d.jpg!cr_1125x549_0_72!q70.jpg.dpg",
          url: ""
        },
        {
          img:
            "http://m.360buyimg.com/mobilecms/s700x280_jfs/t1/53906/16/8649/132967/5d60e97fE0531f445/d718d733c1082477.jpg!cr_1125x549_0_72!q70.jpg.dpg",
          url: ""
        },
        {
          img:
            "http://m.360buyimg.com/mobilecms/s700x280_jfs/t1/50513/38/6391/140522/5d412c4dE8148b43b/ccad1f6b3193e70f.jpg!cr_1125x549_0_72!q70.jpg.dpg",
          url: ""
        }
      ]
    };
  },
  methods: {
    changePage(current) {
      console.log("当前轮播图序号为:" + current);
    },
    clickHandler(item, index) {
      console.log(item, index);
    }
  }
};
</script>

<style scoped>
@import "index.css";
</style>