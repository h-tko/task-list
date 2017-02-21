<template>
    <transition name="fade" tag="div" v-if="show">
        <form>
            <div class="card mt-20">
                <div class="card-header">要望登録</div>
                <div class="card-block">
                    <div class="row">
                        <div class="col">
                            <label for="Title">タイトル</label>
                            <input type="text" name="Title" id="Title" class="form-control" v-model="title">
                        </div>
                    </div>
                    <div class="row">
                        <div class="col">
                            <label for="Title">本文</label>
                            <textarea class="form-control" rows="6" v-model="body"></textarea>
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
            body: null,
            show: true,
        }
    },
    methods: {
        regist() {
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
