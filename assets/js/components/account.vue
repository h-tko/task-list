<template>
    <div class="mt-20">
        <ol class="breadcrumb">
            <li class="breadcrumb-item"><router-link to="/tasks">要望一覧</router-link></li>
            <li class="breadcrumb-item active">アカウント登録</li>
        </ol>

        <transition name="custom-classes-transition" enter-active-class="animated fadeInLeft" leave-active-class="animated fadeOut" mode="out-in" appear>
            <div class="mt-5 card" v-if="show">
                <div class="card-block">
                    <h4 class="card-title text-primary">新規アカウント登録</h4>
                    <p class="card-text">登録するアカウント情報を入力してください。</p>

                    <div class="row mt-3">
                        <div class="col">
                            <div class="input-group">
                                <span class="input-group-addon">メールアドレス</span>
                                <input type="text" v-model="mail_address">
                            </div>
                        </div>
                    </div>

                    <div class="row mt-3">
                        <div class="col">
                            <div class="input-group">
                                <span class="input-group-addon">名前</span>
                                <input type="text" v-model="name">
                            </div>
                        </div>
                    </div>

                    <div class="row mt-3">
                        <div class="col">
                            <div class="input-group">
                                <span class="input-group-addon">パスワード</span>
                                <input type="password" v-model="password">
                            </div>
                        </div>
                    </div>

                    <div class="row mt-3 mr-auto">
                        <div class="col">
                            <button type="button" class="btn btn-outline-primary" v-on:click="newAccount()">登録</button>
                        </div>
                    </div>
                </div>
            </div>
        </transition>
    </div>
</template>
<script>

export default {
    name: 'Account',
    data() {
        setTimeout(() => {
            this.show = true
        }, 100);
        return {
            show: false,
            mail_address: null,
            name: null,
            password: null,
        }
    },
    methods: {
        newAccount() {

            $.post('/account/new', {
                mail_address: this.mail_address,
                name: this.name,
                password: this.password
            }, (result) => {
                if (result.err) {
                    console.log(err)
                } else {
                    this.$router.push('/tasks')
                }
            })

            this.show = false
            this.$router.push('/tasks')
        },
    }
}
</script>
