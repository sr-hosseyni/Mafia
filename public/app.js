new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        message: {
            status: false,
            message: '',

            // game concepts
            isTrue: true,
            isNight: false,
            isFinished: false,
            aliveCitizens: [],
        },
        username: null,
        password: null,
        joined: false // True if email and username have been filled in
    },

    created: async function () {
        this.connect();
        while (!this.ws.readyState) {
            await new Promise(r => setTimeout(r, 1000));
        }
        if (this.username = localStorage.getItem('username')) {
            console.log(this.username)
            this.joined = true;
            this.login();
        }
    },

    methods: {
        login: function () {
            this.ws.send(
                JSON.stringify({
                        name: this.username,
                        password: this.password,
                    }
                ));
        },

        join: function () {
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            // this.username = $('<p>').html(this.username).text();
            this.joined = true;
            localStorage.setItem('username', this.username);
            this.login();
        },

        connect: function () {
            this.ws = new WebSocket('ws://' + window.location.host + '/ws');
            this.ws.addEventListener('message', this.receive);
        },

        receive: function (e) {
            this.message = JSON.parse(e.data);
        },

        gravatarURL: function (email) {
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        }
    }
});