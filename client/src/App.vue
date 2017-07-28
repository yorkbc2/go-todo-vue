<template>
  <div id="app">
    <nav class="navbar navbar navbar-default navbar-fixed-top">
    <div class="container-fluid">
      <div class="navbar-header">
        <router-link class="navbar-brand" to="/">YourTODO</router-link>
      </div>
      <ul class="nav navbar-nav">
        <li><router-link to="/">Home</router-link></li>
      </ul>
      <ul class="nav navbar-nav navbar-right">
      <li v-if="!user"><router-link to="/signup"><span class="glyphicon glyphicon-user"></span> Sign Up</router-link></li>
      <li v-if="!user"><router-link to="/signin"><span class="glyphicon glyphicon-log-in"></span> Login</router-link></li>
      <li v-if="user">
        <router-link to="/profile"><span class="glyphicon glyphicon-user"></span> Profile</router-link></li>
      <li v-if="user">
        <router-link to="/logout">
          <span class="glyphicon glyphicon-log-in"></span>
          Logout
        </router-link>
      </li>
    </ul>
    </div>
  </nav>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'app',
  data: function (){
    return {
      user: null
    }
  },
  mounted() {
    if(typeof window.sessionStorage.getItem('user') == 'string') {
      if (/^[\],:{}\s]*$/.test(window.sessionStorage.getItem("user").replace(/\\["\\\/bfnrtu]/g, '@').
      replace(/"[^"\\\n\r]*"|true|false|null|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?/g, ']').
      replace(/(?:^|:|,)(?:\s*\[)+/g, ''))) {
        this.user = JSON.parse(window.sessionStorage.getItem("user"))
      }
      else {
        this.user = window.sessionStorage.getItem("user")
      }
    }
    else {
      this.user = null
    }

  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 35px;
}
</style>
