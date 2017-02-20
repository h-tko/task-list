<template>
    <div class="modal fade" id="login-modal">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">ログイン</h5>
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                </div>
                <div class="modal-body">
                    <div class="row">
                        <div class="col">
                            <div class="input-group">
                                <span class="input-group-addon">E-Mail</span>
                                <input type="text" class="form-control" v-model="mail_address">
                            </div>
                        </div>
                    </div>
                    <div class="row mt-20">
                        <div class="col">
                            <div class="input-group">
                                <span class="input-group-addon">PASSWORD</span>
                                <input type="password" class="form-control" v-model="password">
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-primary" v-on:click="login()">Login</button>
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'LoginModal',
    data() {
        return {
            mail_address: null,
            password: null,
            memberID: null,
        }
    },
    methods: {
        login() {
            $.post("/login", {mail_address: this.mail_address, password: this.password}, (result) => {
                if (result.err) {
                    alert(result.err)
                } else {
                    this.memberID = result.MemberID
                }
            })
        }
    },
    watch: {
        memberID: (newMemberID) => {
            window.MemberID = newMemberID

            const objects = document.querySelectorAll(".login-disabled")
            const enableObjects = document.querySelectorAll(".login-enabled")

            objects.forEach((val) => {
                val.style.display = 'none'
            })

            enableObjects.forEach((val) => {
                val.style.display = 'block'
            })

            $('#login-modal').modal('hide')
        }
    }
}
</script>
