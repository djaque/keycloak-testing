<template>
	<form @submit.prevent="handleSubmit">
		<h3>Registro</h3>

		<div class="form-group">
			<label for="">Nombre</label>
			<input type="text" name="" id="" class="form-control" placeholder="Nombre" v-model="first_name">
		</div>	

		<div class="form-group">
			<label for="">Apellido</label>
			<input type="text" name="" id="" class="form-control" placeholder="Apellido" v-model="last_name">
		</div>	
		
		<div class="form-group">
			<label for="">Email</label>
			<input type="email" name="" id="" class="form-control" placeholder="Email" v-model="email">
		</div>	

		<div class="form-group">
			<label for="">Contraseña</label>
			<input type="password" name="" id="" class="form-control" placeholder="Contraseña" v-model="password">
		</div>	

		<div class="form-group">
			<label for="">Confirma tu Contraseña</label>
			<input type="password" name="" id="" class="form-control" placeholder="Contraseña" v-model="password_confirm">
		</div>		

		<button class="btn btn-primary btn-block" >Registrar</button>	
	</form>
</template>
 
<script>

import axios from 'axios'

export default {
	name: 'Register',
	data() {
		return {
			first_name:'',
			last_name:'',
			email:'',
			password:'',
			password_confirm:'',
		}
	},
	methods: {
		async handleSubmit() {
			await axios.post("register", {
				first_name: this.first_name,
				last_name: this.last_name,
				email: this.email,
				password: this.password,
				password_confirm: this.password_confirm,
			});
			this.$router.push('/login');
		}
	}
}
</script>