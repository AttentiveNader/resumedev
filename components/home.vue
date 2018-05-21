<template>
	<div class="home">
		<img class="background-img" src="~/static/white-background.jpg">
		<h1>Nader Atef</h1>
		<div class="bio">
			<div class="sub-bio">
				<p>
					I'm a self-taught Web developer
					I'm really passionate about programming in general and learning new technologies
					now I'm trying to get more experience as a fullstack web developer and to dive into machine learning 
				</p>
			</div>
		</div>
		<div class="skills-cont">
			<div class="frontend">
				<h3>FrontEnd:</h3>
				<ul>
					<li v-bind:key="skill" v-for="skill in skills.frontend">
						<p>{{skill}}</p>
					</li>
				</ul>
			</div>
			<div class="backend">
				<h3>BackEnd:</h3>
				<ul>
					<li v-bind:key="skill" v-for="skill in skills.backend">
						<p>{{skill}}</p>
					</li>
				</ul>
			</div>
		</div>
	</div>
</template>
<script>

export default {
	data(){
		return {
			skills :{
				frontend :["HTML", "CSS" , "Javascript", "Vuejs/Nuxt","Responsive Design","XML & XSL","jQuery..."],
				backend : ["Golang","MySQL"]
			}
		}
	}
}
	
</script>
<style type="text/css" scoped>
	.background-img{
		position: fixed;
		top: 0;
		left: 0;
		height: 100%;
		z-index: -5;
		width: 100%;
	}
	.home{
		padding:20px;
	}
	.sub-bio{
		padding:5px;
		margin: 0px auto;
		max-width: 600px;
		text-align: center;
		font-size: larger;
		font-weight: medium;
		background: white ;
		background: rgba(255,255,255,0.3);
	}
	.skills-cont{
		display: flow-root;
		border-radius: 12px;
		width: 100%;
		margin:0 auto;
		max-width: 600px;
		padding:10px;
		background: white ;
		background: rgba(255,255,255,0.3);
		-webkit-box-sizing: border-box;
		        box-sizing: border-box;
	}
	 @media screen and (max-width: 600px){
	 	.skills-cont{
	 		width: 100%;
	 		max-width: none;
	 	}
	 }
	.skills-cont div{
		float:left;
		text-align: left;
		display: inherit;
		width: 50%;
		font-size: larger;
		font-weight: bolder;
		-webkit-box-sizing: border-box;
		        box-sizing: border-box;
	}
	h1{		
		margin-bottom: 10px;
	}
</style>