<template>
    <div class="mt-20">
        <ol class="breadcrumb" v-if="show">
            <li class="breadcrumb-item"><router-link to="/tasks">要望一覧</router-link></li>
            <li class="breadcrumb-item active">{{task.Title}}</li>
        </ol>
        <transition name="custom-classes-transition" enter-active-class="animated fadeIn" v-on:enter="enter" mode="out-in" tag="div" appear>
            <div v-if="show">
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

                        <transition-group name="custom-classes-transition" enter-active-class="animated slideInLeft" tag="div" appear>
                            <div class="col-sm-8 float-sm-right mt-2" v-for="comment in taskComments" v-bind:key="comment" v-if="showComment">
                                <div class="card card-outline-info">
                                    <div class="card-block">
                                        <p v-html="nl2br(comment.Comment)"></p>
                                        <div class="float-sm-right text-muted"><small>{{comment.Member.Name}}が{{datetimeFormat(comment.CreatedAt)}}に投稿しました</small></div>
                                    </div>
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
                                <div v-if="err" v-for="errObj in err" v-bind:key="errObj">
                                    <span class="text-error" v-if="invalidComment(errObj)">{{getErrorMessageInComment(errObj)}}</span>
                                </div>
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
            showComment: false,
            comment: null,
            memberID: null,
            member: null,
            err: null,
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
                    vm.showComment = false
                    vm.memberID = result.MemberID
                    vm.member = result.Member
                })
            }
        })
    },
    methods: {
        nl2br(value) {
            return this.$options.filters.nl2br(value)
        },
        datetimeFormat(value) {
            return this.$options.filters.datetime_format(value)
        },
        logined() {
            return this.memberID !== null && parseInt(this.memberID) > 0
        },
        invalidComment(err) {
            return err && err.Field === "Comment"
        },
        getErrorMessageInComment(err) {
            if (err.Kind === "required") {
                return "コメントを入力してください。"
            }
        },
        clearComment() {
            this.comment = null
        },
        enter(el) {
            if (this.taskComments !== null) {
                this.showComment = true
            }
        },
        sendComment() {
            $.post('/detail/send_comment/', {
                id: this.task.ID,
                comment: this.comment
            }, (result) => {
                if (result.err) {
                    this.err = result.err
                } else {
                    const comment = {
                        Comment: this.comment,
                        CreatedAt: new Date().toString(),
                        Member: {
                            Name: this.member.Name
                        }
                    }

                    this.taskComments.push(comment)
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
                    this.showComment = false
                    this.memberID = result.MemberID
                    this.member = result.Member
                }
            })
        },
    }
}
</script>
