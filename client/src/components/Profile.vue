<template>
    <div class="container">
        <h1>
            Your profile
        </h1>
        <div class="col-md-6 col-md-offset-3">
          <div class="panel panel-body">
            <form v-on:submit.prevent="AddTodo($event)" class="form-field">
              <div class="input-group">
                <input v-model="todo.content" type="text" class="form-control" placeholder="Your todo">
                <div class="input-group-btn">
                  <button type="submit" class="btn btn-primary">Add Todo +</button>
                </div>
              </div>
            </form>
            <br>
            <table class="table table-striped">
              <thead>
                <tr>
                  <th width="75%">Todo: </th>
                  <th width="25%">Event: </th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in todos">
                  <td style="text-align: left; line-height: 130%">{{item.Content}}</td>
                  <td><button class="btn btn-danger">Remove</button></td>
                </tr>
              </tbody>
            </table>
            <!-- <ul class="list-group">
              <li v-for="item in todos" class="list-group-item">
                {{item.Content}} <button class="btn btn-danger">Remove &times;</button>
              </li>
            </ul> -->
          </div>
        </div>
    </div>
</template>

<script>
  import config from "./../etc/config.js"
  export default {
    name: "Profile",

    data() {
      return {
        todo: {
          content: ""
        },
        todos: []
      }
    },

    methods: {
      AddTodo(ev) {
        var id = JSON.parse(window.sessionStorage.getItem("user")).UKEY,
          content = this.todo.content;


        this.$http.post(config.baseUrl + "todo/create", {
          ukey: id,
          content: content
        }, {emulateJSON: true}).then(res => {
          this.todos.unshift(res.body)
          this.todo.content = ""
        }, err => console.error(err))

      },

      GetTodos() {
        var id = JSON.parse(window.sessionStorage.getItem("user")).ID

        this.$http.post(config.baseUrl + "todos", {
          id: id
        }, {emulateJSON: true}).then(res => {
          if(res.body.length > 0) {
            this.todos = res.body
          }
          else {
            this.todos = []
          }
        }, err => console.error(err))
      }
    },

    created() {

      var session = window.sessionStorage

      if(JSON.parse(window.sessionStorage.getItem("user")).UKEY.length < 1) {
        this.$router.push('/signup')
      }
      else {
        this.GetTodos()
      }

    }
  }
</script>

<style scoped>
  li {
    text-align: left!important;
  }
</style>
