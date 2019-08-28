import axios from 'axios';
import Qs from 'qs';
// import { promised } from 'q';
axios.defaults.withCredentials = true;
axios.defaults.baseURL = 'http://localhost';
axios.defaults.headers.post['Content-Type'] =
  'application/x-www-form-urlencoded;charset=UTF-8';
axios.defaults.transformRequest = [
  function(data, config) {
    // 对 data 进行任意转换处理
    if (config['Content-Type'] == 'multipart/form-data') {
      return data;
    }
    return Qs.stringify(data);
  }
];
// axios.defaults.paramsSerializer = function(params) {
//   params['wxbrowser'] = 'company';
//   return Qs.stringify(params, { arrayFormat: 'brackets' });
// };
// let ajax = axios.create({
//   baseURL: ,
//   timeout: 3000,
//   transformRequest: ,
//   withCredentials: true,
//   headers: {
//     "X-Requested-With": "XMLHttpRequest"
//   }
// });

export function get(url, params = {}) {
  return new Promise((resolve, reject) => {
    axios
      .get(url, {
        params: params
      })
      .then(resp => {
        resolve(resp.data);
        // if (resp.data.status == 962) {
        //   this.$router.push('/login?url=' + this.$route.path);
        //   return;
        // }
      })
      .catch(err => {
        reject(err);
      });
  });
}

export function post(url, data = {}) {
  return new Promise((resolve, reject) => {
    axios
      .post(url, data)
      .then(resp => {
        resolve(resp);
      })
      .catch(err => {
        reject(err);
      });
  });
}

let upvideoconfig = axios.create({
  baseURL: 'http://localhost',
  headers: {
    'Content-Type': 'multipart/form-data'
  }
});
export function upvideo(url, data) {
  return new Promise((resolve, reject) => {
    upvideoconfig
      .post(url, data)
      .then(resp => {
        resolve(resp);
      })
      .catch(err => {
        reject(err);
      });
  });
}
