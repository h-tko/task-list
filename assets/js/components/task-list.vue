<template>
    <div name="fade" class="mt-20">
        <legend>要望一覧</legend>
        <div class="row">
            <div class="col">
                <input class="form-control mr-sm-2" type="text" placeholder="Search" v-model="freeword">
            </div>
            <div class="col">
                <button class="btn btn-outline-success my-2 my-sm-0" type="button" v-on:click="search()">Search</button>
            </div>
        </div>

        <transition-group class="list-group mt-20" name="slider" tag="ul" appear>
            <div v-for="task in tasks" v-bind:key="task">
                <router-link :to="url(task.ID)" exact>
                    <li class="list-group-item justify-content-between">
                        {{task.Title}}
                        <span class="badge badge-default badge-pill">{{task.RegistMember.Name}}</span>
                    </li>
                </router-link>
            </div>
        </transition-group>
    </div>
</template>

<script>
export default {
    data() {
        return {
            tasks: null,
            error: null,
            freeword: null,
            show: true,
        }
    },
    beforeRouteEnter(route, redirect, next) {
        $.get(route.path, (result) => {
            if (result.err) {
                console.log(result.err)
            } else {
                next(vm => {
                    vm.tasks = result.tasks
                })
            }
        })
    },
    methods: {
        url(id) {
            return "/detail/" + id
        },
        search() {
            $.post("/search", {freeword: this.freeword}, (result) => {
                if (result.err) {
                    console.log(result.err)
                } else {
                    this.tasks = result.tasks
                }
            });
        }
    },
    watch: {
        $route () {
            this.tasks = null
            $.get(this.$route.path, (result) => {
                if (result.err) {
                    this.error = result.err.toString()
                } else {
                    this.tasks = result.tasks
                }
            })
        }
    }
}
</script>
