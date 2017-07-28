export default {

  storage: window.sessionStorage,

  log() {
    console.log(this.storage)
    return this.storage
  },

  len() {
    return this.storage.length
  },

  set(key, value) {
    if(key == undefined) {
      throw new Error("[vue-session] :: Key is required")
    }


    if(value == undefined) {
      throw new Error("[vue-session] :: Value is required")
    }

    if(typeof value == "object") {
      value = JSON.stringify(value)
    }

    this.storage.setItem(key ,value)

    return this.storage
  },

  get(key) {

    if(key == undefined) {
      throw new Error("[vue-session] :: Key is required")
    }

    var value = this.storage.getItem(key)

    if (/^[\],:{}\s]*$/.test(value.replace(/\\["\\\/bfnrtu]/g, '@').
    replace(/"[^"\\\n\r]*"|true|false|null|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?/g, ']').
    replace(/(?:^|:|,)(?:\s*\[)+/g, ''))) {
      value = JSON.parse(value)
    }

    return value
  },

  remove(key) {
    this.storage.remove(key)

    return this.storage
  },

  key(num) {
    return this.storage.key(num)
  },

  removeAll() {
    this.storage.clear()
    return this.storage
  },

  isset(key) {
    if(typeof this.get(key) !== 'undefined') {
      return true
    }
    else {
      return false;
    }
  },

  install(Vue,options) {
    Vue.prototype.$session = this

    return Vue
  }}

