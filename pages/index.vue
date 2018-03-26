<template>
  <div id="app">
    <contact></contact>
    <home></home>
    <div id="container"></div>
  </div>
</template>

<script>
import contact from "~/components/contact.vue"
import home from "~/components/home.vue"
export default {
  name: 'App',
  components:{
    contact,
    home
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
  margin-top: 60px;
}
#container {
    display: none;
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
    height: 60px;
    width:60px;
    vertical-align: middle;
    border-radius: 50%;
    font-size: 20px;
    cursor: pointer;
    text-align: center;
  }
  .expanded .pull-btn{
    transform: rotate(180deg);
    -webkit-transform: rotate(180deg);
    -o-transform: rotate(180deg);
    -ms-transform: rotate(180deg);
    -moz-transform: rotate(180deg); 
  }
  .pull-btn span{
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
    transition: color 0.3s ease-in-out, background-color 0.15s ease-in-out,border-color 0.15s ease-in-out, box-shadow 0.3s ease-in-out;
    cursor: pointer;
    outline: none;

  }
  button:active{
    box-shadow: 0 0 3px rgba(100, 250, 200, 1) !important;
  }
</style>
