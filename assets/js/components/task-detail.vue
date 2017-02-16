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
                        <div class="card card-outline-primary mt-4">
                            <div class="card-header card-primary">{{task.Title}}</div>
                            <div class="card-block">
                                <p class="card-text" v-html="nl2br(task.Body)"></p>

                                <div class="float-sm-right">
                                    <p><small><strong>要望者:&nbsp;</strong>{{task.RegistMember.Name}}</small></p>
                                </div>
                            </div>
                        </div>

                        <transition-group name="custom-classes-transition" enter-active-class="animated slideInLeft" appear>
                            <div class="col-sm-8 float-sm-right mt-20" v-for="comment in taskComments" v-bind:key="comment" v-if="taskComments">
                                <div class="card card-outline-success mt-2">
                                    <div class="card-block" v-html="nl2br(comment.Comment)"></div>
                                </div>
                            </div>
                        </transition-group>
                    </div>
                </div>

                <div class="row mt-5" v-if="logined()">
                    <div class="col">
                        <div class="card">
                            <div class="card-block">
                                <textarea class="form-control" v-model="comment" rows="6" placeholder="コメントを書く"></textarea>
                            </div>
                            <div class="card-footer">
                                <div class="row">
                                    <div class="col-sm-2 float-sm-right">
                                        <button type="button" class="btn btn-outline-default" v-on:click="clearComment()">クリア</button>
                                        <button type="button" class="btn btn-outline-primary" v-on:click="sendComment()">投稿</button>
                                    </div>
                                </div>
                            </div>
                        </div>
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
            taskComments: null,
            show: false,
            comment: null,
            memberID: null,
        }
    },
    beforeRouteEnter(route, redirect, next) {
        $.get(route.path, (result) => {
            if (!result) {
                console.log("err!")
            } else {
                next(vm => {
                    vm.task = result.Task
                    vm.taskComments = result.TaskComments
                    vm.show = true
                    vm.memberID = result.MemberID
                })
            }
        })
    },
    methods: {
        nl2br(value) {
            return this.$options.filters.nl2br(value)
        },
        logined() {
            return this.memberID !== null
        },
        clearComment() {
            this.comment = null
        },
        sendComment() {
            var comment = {
                Comment: this.comment
            }
            this.taskComments.push(comment)

            $.post('/detail/send_comment/', {
                id: this.task.ID,
                comment: this.comment
            }, (result) => {
                if (!result) {
                    connsole.log(result)
                } else {
                    this.comment = null
                }
            })
        },
    },
    watch: {
        $route() {
            this.task = null
            this.taskComments = null
            $.get(this.$route.path, (result) => {
                if (!result) {
                    console.log("err!")
                } else {
                    this.task = result.Task
                    this.taskComments = result.TaskComments
                    this.show = true
                    this.memberID = result.MemberID
                }
            })
        },
    }
}
</script>
