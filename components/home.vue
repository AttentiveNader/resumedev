<template>
	<div class="home">
		<img class="background-img" src="~/static/white-background.jpg">
		<h1>Nader Atef</h1>
		<div class="skills-cont">
			<div class="frontend">
				<h3>Front-End :</h3>
				<ul>
					<li v-bind:key="skill" v-for="skill in skills.frontend">
						<p>{{skill}}</p>
					</li>
				</ul>
			</div>
			<div class="backend">
				<h3>Back-End :</h3>
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
				frontend :["HTML", "CSS" , "Javascript", "Vuejs/Nuxt","Responsive Design","XML & XSL"],
				backend : ["Golang"]
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
	.skills-cont{
		margin:0 auto;
		max-width: 600px;
	}
	.skills-cont div{
		float:left;
		text-align: left;
	}
	h1{		
		margin-bottom: 20px;
	}
</style>