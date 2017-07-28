<template>
  <div class="container-fluid">
    <div id="myCarousel" class="carousel slide" data-ride="carousel" style="max-height: 640px!important">
      <!-- Indicators -->
      <ol class="carousel-indicators">
        <li data-target="#myCarousel" data-slide-to="0" class="active"></li>
        <li data-target="#myCarousel" data-slide-to="1"></li>
        <li data-target="#myCarousel" data-slide-to="2"></li>
      </ol>

      <!-- Wrapper for slides -->
      <div class="carousel-inner">
        <div class="item active">
          <img src="http://www.officelovin.com/wp-content/uploads/2016/01/ragged-edge-london-office-6.jpg" alt="Los Angeles" style="max-height: 600px!important; width: 100%;" height="640px">
          <div class="carousel-caption">
            <h1>Make your work better</h1>
            <router-link to="/signup" class="btn btn-success btn-fluid">
              SIGN UP FOR FREE
            </router-link>
          </div>
        </div>

        <div class="item">
          <img src="https://edge.alluremedia.com.au/m/l/2016/02/3-office-space.jpg" style="max-height: 600px!important; width: 100%;" alt="Chicago" height="640px">
          <div class="carousel-caption">
            <h1>Make your work better</h1>
            <router-link to="/signup" class="btn btn-success btn-fluid">
              SIGN UP FOR FREE
            </router-link>
          </div>

        </div>

        <div class="item">
          <img src="http://maintainwithfresh.com/images/office/office2.jpg" style="max-height: 600px!important; width: 100%;" alt="New York" height="640px">
          <div class="carousel-caption">
            <h1>Make your work better</h1>
            <router-link to="/signup" class="btn btn-success btn-fluid">
              SIGN UP FOR FREE
            </router-link>
          </div>
        </div>
      </div>

      <!-- Left and right controls -->
      <a class="left carousel-control" href="#myCarousel" data-slide="prev">
        <span class="glyphicon glyphicon-chevron-left"></span>
        <span class="sr-only">Previous</span>
      </a>
      <a class="right carousel-control" href="#myCarousel" data-slide="next">
        <span class="glyphicon glyphicon-chevron-right"></span>
        <span class="sr-only">Next</span>
      </a>
    </div>
    <div class="cointaner about">
      <div class="row">
        <div class="col-md-4">
          <span class="glyphicon glyphicon-user"></span>
          <h3>
            Free profiles
          </h3>
          <p>
            Register for free now and get profile for you. If you sign up, you will have big permissions
          </p>
          <p>
            <router-link to="/signup" class="btn btn-primary btn-fluid">
              SIGN UP FOR FREE
            </router-link>
          </p>
        </div>
        <div class="col-md-4">
          <span class="glyphicon glyphicon-th-list"></span>
          <h3>
            Powerful lists
          </h3>
          <p>
            Create powerful lists with your todos for free
          </p>
        </div>
        <div class="col-md-4">
          <span class="glyphicon glyphicon-eye-open"></span>
          <h3>
            Easy design
          </h3>
          <p>
            Design of our TODO LIST is very good for noobs
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    mounted() {
    }
  }
</script>

<style>
  .about {
    padding: 40px 0;
  }
  .about .row .col-md-4 span {
    font-size: 36pt!important;
    color: #ff3333;
  }
</style>
