<template>

  <div class="container">
    <h1>
      Sign up for free
    </h1>
    <div class="row">
        <div class="col-md-4 col-md-offset-4">
            <div class="panel panel-default">
              <div class="panel-body">
                <div v-if="error" class="alert alert-danger">
                    {{this.error}}
                </div>
                <form v-on:submit.prevent="SignUpForm($event)" class="form-field">
                  <div class="input-group">
                    <span class="input-group-addon"><i class="glyphicon glyphicon-user"></i></span>
                    <input required type="text" class="form-control" v-model="user.login" placeholder="Your login">
                  </div>
                  <br>
                  <div class="input-group">
                    <span class="input-group-addon"><i class="glyphicon glyphicon-lock"></i></span>
                    <input required type="password" class="form-control" v-model="user.password" placeholder="Your password">
                  </div>
                  <br>
                  <div class="input-group">
                    <span class="input-group-addon"><i class="glyphicon glyphicon-lock"></i></span>
                    <input required type="password" class="form-control" v-model="check.password" placeholder="Retry your password">
                  </div>
                  <br>
                  <div>
                    <button type="submit" id="submitButton" class="btn btn-success">
                        Sign up
                    </button>
                    <button type="reset" class="btn btn-danger">
                        Clear
                    </button>
                  </div>
                </form>
              </div>
            </div>
        </div>
    </div>
  </div>

</template>

<script>
    import config from "./../etc/config.js"

  export default {
    name: "SignUp",

    data: function () {
      return {
        baseAPIUrl: config.baseUrl,
        user: {},
        check: {},
        error: ""
      }
    },

    methods: {
      SignUpForm(ev) {
          if(this.user.password == this.check.password) {
              this.$http.post(this.baseAPIUrl + "signup", this.user, {emulateJSON: true})
                  .then(res => {
                    if(typeof res.body.UKEY == "string") {
                      if(res.body.UKEY.length > 1) {

                        window.sessionStorage.setItem("user", res.bodyText)

                        this.$router.push("/profile")
                      }
                    }
                    else {
                      this.error = "This login has already used"
                    }

                  }, err => console.error(err))
          }
          else {
              this.error = "Password must be equality"
          }
      }
    },

    created() {

      var session = window.sessionStorage
      if(typeof window.sessionStorage.getItem("user") == "string")
      {
        if(typeof JSON.parse(window.sessionStorage.getItem("user")).UKEY == "string") {
          if(JSON.parse(window.sessionStorage.getItem("user")).UKEY.length > 1) {
            this.$router.push('/profile')
          }
        }
    }


    }
  }
</script>
