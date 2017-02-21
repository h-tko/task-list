<template>
    <transition name="fade" tag="div" v-if="show">
        <form>
            <div class="card mt-20">
                <div class="card-header">要望登録</div>
                <div class="card-block">
                    <div class="row">
                        <div class="col">
                            <div class="form-group">
                                <label for="Title">タイトル</label>
                                <input type="text" name="Title" id="Title" class="form-control" v-model="title">
                            </div>
                            <span class="text-danger" v-if="alert_title">タイトルが正しく入力されていません</span>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col">
                            <div class="form-group">
                                <label for="Body">本文</label>
                                <textarea class="form-control" rows="6" id="Body" v-model="body"></textarea>
                            </div>
                            <span class="text-danger" v-if="alert_body">本文が正しく入力されていません</span>
                        </div>
                    </div>
                    <div class="row mt-20">
                        <div class="col">
                            <button type="button" class="btn btn-primary" v-on:click="regist">登録</button>
                        </div>
                    </div>
                </div>
            </div>
        </form>
    </transition>
</template>

<script>
export default {
    data() {
        return {
            title: null,
            alert_title: false,
            body: null,
            alert_body: false,
            show: true,
        }
    },
    methods: {
        regist() {

            this.alert_title = (this.title == null || this.title.length < 1)
            this.alert_body = (this.body == null || this.body.length < 1)

            if (this.alert_title || this.alert_body) {
                return
            }

            $.post("/regist_task", {title: this.title, body: this.body}, (result) => {
                if (result.status !== "ok") {
                    console.log("error!")
                } else {
                    alert('登録しました。')

                    this.show = false

                    this.$router.push('/')
                }
            });
        }
    }
}
</script>
