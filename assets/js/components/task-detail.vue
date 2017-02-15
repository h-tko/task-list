<template>
    <div class="mt-20">
        <ol class="breadcrumb" v-if="show">
            <li class="breadcrumb-item"><router-link to="/tasks">要望一覧</li>
            <li class="breadcrumb-item active">{{task.Title}}</li>
        </ol>
        <transition name="fade" mode="out-in" tag="div" v-if="show">
            <div>
                <div class="row">
                    <div class="col">
                        <div class="card card-outline-primary mt-20">
                            <div class="card-header card-primary">{{task.Title}}</div>
                            <div class="card-block">
                                <p class="card-text" v-html="nl2br(task.Body)"></p>
                            </div>
                        </div>

                        <div class="col-sm-8 float-sm-right mt-20">
                            <div class="card card-outline-success mt-3">
                                <div class="card-block">
                                    ほげほげほげほ
                                </div>
                            </div>
                        </div>
                        <div class="col-sm-8 float-sm-right">
                            <div class="card card-outline-success mt-3">
                                <div class="card-block">
                                    ほげほげほげほ
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="row mt-5">
                    <div class="col-sm-2 float-sm-right">
                        <button type="button" class="btn btn-outline-primary">コメントを書く</button>
                    </div>
                </div>
            </div>
        </transition>
    </div>
</template>

<script>
export default {
    name: 'TaskDetail',
    data() {
        return {
            task: null,
            show: false,
        }
    },
    beforeRouteEnter(route, redirect, next) {
        $.get(route.path, (result) => {
            if (!result) {
                console.log("err!")
            } else {
                next(vm => {
                    vm.task = result.Task
                    vm.show = true
                })
            }
        })
    },
    methods: {
        nl2br(value) {
            return this.$options.filters.nl2br(value)
        }
    },
    watch: {
        $route() {
            this.task = null
            $.get(this.$route.path, (result) => {
                if (!result) {
                    console.log("err!")
                } else {
                    this.task = result.Task
                    this.show = true
                }
            })
        }
    }
}
</script>
