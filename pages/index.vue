<template>
  <div id="app">
    <contact></contact>
    <home></home>
    <portfolio></portfolio>
    <div id="container"></div>
    <div id="loading">
      <div class="circle"></div>
    </div>
  </div>
</template>

<script>
import contact from "~/components/contact.vue"
import portfolio from "~/components/portfolio.vue"
import home from "~/components/home.vue"
export default {
  name: 'App',
  components:{
    contact,
    home,
    portfolio
  }
}
</script>

<style>
#app {
  font-family:  'Ubuntu Condensed','Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  font-size: larger
}
p{
  padding :2px ;
}
a{
  text-decoration: none;
  color : black ;
  font-weight: bolder;
}
#container {
    display: none;
    opacity: 0 ;
    -webkit-transition: 0.3s all;
    -o-transition: 0.3s all;
    transition: 0.3s all;
    position: fixed;
    left:0;
    top:0;
    height:100%;
    width:100%;
    z-index: 4;
    background:rgba(0,0,0,0.4);
  }
.rotated{
    transform: rotate(180deg);
    -webkit-transform: rotate(180deg);
    -o-transform: rotate(180deg);
    -ms-transform: rotate(180deg);
    -moz-transform: rotate(180deg)
}
  .pull-btn{
    height: 40px;
    width:40px; 
    font-size: 30px !important;
    vertical-align: middle;
    border-radius: 50%;
    font-size: 16px;
    cursor: pointer;
    border-radius: 50% ;
    color: aliceblue ;
    text-align: center;
    background: #2c3e50 ;
    -webkit-transition: 0.5s ;
    -o-transition: 0.5s ;
    transition: 0.5s ;
    text-align: center;
    
  }
 
  #loading{
  display:none; 
  position: fixed;
  top:0;
  left:0;
  z-index: 12;
  width: 100%;
  height:100%;
  background: #3d3d3d;
  opacity: 0.6;
}
.circle{
  height:100px;
  width: 100px;
  z-index: 10;
  position: absolute;
  top:calc(50% - 50px);
  left:calc(50% - 50px);
  opacity:1.4;
  border-radius: 50%;
  border-top:10px solid black;
  border-left:10px solid #c53f55;
  border-bottom:10px solid white;
  border-right: none;
  -webkit-animation: 0.4s rotate infinite ease-in-out ;
          animation: 0.4s rotate infinite ease-in-out ;
}
@-webkit-keyframes rotate {
  0%{
    -webkit-transform: rotate(0deg);
            transform: rotate(0deg); 
  }
  50%{
    -webkit-transform: rotate(160deg);
            transform: rotate(160deg);
  }
  100%{
    -webkit-transform: rotate(360deg);
            transform: rotate(360deg);
  }
}

@keyframes rotate {
  0%{
    -webkit-transform: rotate(0deg);
            transform: rotate(0deg); 
  }
  50%{
    -webkit-transform: rotate(160deg);
            transform: rotate(160deg);
  }
  100%{
    -webkit-transform: rotate(360deg);
            transform: rotate(360deg);
  }
}
   .pull-btn:hover{
        -webkit-box-shadow: 0 0 6px black,0 0  6px black ,0 0  6px black ,0 0  6px black ;
                box-shadow: 0 0 6px black,0 0  6px black ,0 0  6px black ,0 0  6px black ;
  }

  .pull-btn:active{
    -webkit-box-shadow: 0 0 6px black,0 0  6px black ,0 0  6px black ,0 0  6px black ;
            box-shadow: 0 0 6px black,0 0  6px black ,0 0  6px black ,0 0  6px black ;
  }
  .pull-btn span{
    -webkit-transition: 0.3s;
    -o-transition: 0.3s;
    transition: 0.3s;
    vertical-align: middle;
  }
  button{
    display: inline-block;
    font-weight: 400;
    text-align: center;
    white-space: nowrap;
    vertical-align: middle;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    border: 1px solid transparent;
    padding: 0.375rem 0.75rem;
    font-size: 1rem;
    line-height: 1.5;
    border-radius: 0.25rem;
    -webkit-transition: color 0.3s ease-in-out, background-color 0.15s ease-in-out,border-color 0.15s ease-in-out, -webkit-box-shadow 0.3s ease-in-out;
    transition: color 0.3s ease-in-out, background-color 0.15s ease-in-out,border-color 0.15s ease-in-out, -webkit-box-shadow 0.3s ease-in-out;
    -o-transition: color 0.3s ease-in-out, background-color 0.15s ease-in-out,border-color 0.15s ease-in-out, box-shadow 0.3s ease-in-out;
    transition: color 0.3s ease-in-out, background-color 0.15s ease-in-out,border-color 0.15s ease-in-out, box-shadow 0.3s ease-in-out;
    transition: color 0.3s ease-in-out, background-color 0.15s ease-in-out,border-color 0.15s ease-in-out, box-shadow 0.3s ease-in-out, -webkit-box-shadow 0.3s ease-in-out;
    cursor: pointer;
    outline: none;

  }
  button:active{
    -webkit-box-shadow: 0 0 3px rgba(100, 250, 200, 1) !important;
            box-shadow: 0 0 3px rgba(100, 250, 200, 1) !important;
  }
</style>
