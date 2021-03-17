
const HOME_EVENT_DEVICE_UPDATE = "HOME_EVENT_DEVICE_UPDATE"
const HOME_EVENT_SONG_UPDATE = "HOME_EVENT_SONG_UPDATE"
const HOME_EVENT_DEVICE_LOGIN  = "HOME_EVENT_DEVICE_LOGIN"
const HOME_EVENT_DEVICE_LOGOUT = "HOME_EVENT_DEVICE_LOGOUT"
const HOME_EVENT_MUSIC_PLAY = "HOME_EVENT_MUSIC_PLAY"
const HOME_EVENT_VOLUME_CHANGE = "HOME_EVENT_VOLUME_CHANGE"
const HOME_EVENT_MUSIC_STOP = "HOME_EVENT_MUSIC_STOP"

new Vue({
    el: '#app',
    data() {
        return {
            is_online : true,
            ws : null,
            user : {},
            media : {
                audio_player: null,
                current_device : null,
                volume : 5,
                current_song: {
                    id: "",
                    user_id: "",
                    title: "",
                    description: "",
                    file_path: ""
                }
            },
            device : {
                devices : [],
                query : {
                    filters: {
                        user_id : ""
                    },
                    search: {},
                    orders: {
                        name: "ASC"
                    },
                    offset: 0,
                    limit: 10
                },
            },
            music : {
                musics : [],
                query : {
                    filters: {
                        user_id : ""
                    },
                    search: {},
                    orders: {
                        title: "ASC"
                    },
                    offset: 0,
                    limit: 10
                },
            },
            host : {
                name : "",
                protocol : "",
                port : ""
            }
        }
    },
    created(){
        window.addEventListener('offline', () => { this.is_online = false })
        window.addEventListener('online', () => { this.is_online = true })
        window.history.pushState({ noBackExitsApp: true }, '')
        window.addEventListener('popstate', this.backPress )
        this.setCurrentHost()
    },
    mounted(){
        this.loadSession()
    },
    methods : {
        loadDevices(){

            this.device.query.filters.user_id = this.user.id

            axios
                .post(this.baseUrl() + '/api/v1/devices-list',this.device.query)
                .then(response => {
                    if (response.data.error != null){
                        return
                    }
                    if (response.data.data == null){
                        return
                    }
                    this.device.devices = response.data.data

                })
                .catch(errors => {
                    console.log(errors)
                }) 
        },
        loadMusics(){

            this.music.query.filters.user_id = this.user.id

            axios
                .post(this.baseUrl() + '/api/v1/musics-list',this.music.query)
                .then(response => {
                    if (response.data.error != null){
                        return
                    }
                    if (response.data.data == null){
                        return
                    }
                    this.music.musics = response.data.data

                })
                .catch(errors => {
                    console.log(errors)
                }) 
        },
        openUploadFile(){
            this.$refs["file-upload"].click()
        },
        onFileChange(f) {
            let files = f.target.files || f.dataTransfer.files
            if (!files.length) return

            if (!window.localStorage.getItem('session')) {
                return;
            }

            for (let e of files) {
                this.uploadSong(e,(path => {
                        let music = {
                            user_id: this.user.id,
                            title: e.name,
                            description: "song with tittle : " + e.name,
                            file_path: path,
                        }
                        this.addMusic(music)
                    })
                )

            }

            this.$refs["file-upload"].value = null
        },
        uploadSong(file,done){
            let formData = new FormData();
            formData.append('file', file);
            axios.post(this.baseUrl() + '/api/v1/upload', formData, {
                headers: {
                'Content-Type': 'multipart/form-data'
                }
            }).then(response => {
                if (response.data.error != null){
                    return
                }
                done(response.data.data.path)
            })
            .catch(errors => {
                console.log(errors)
            }) 
        },
        addMusic(music){
            axios
                .post(this.baseUrl() + '/api/v1/musics',music)
                .then(response => {
                    if (response.data.error != null){
                        return
                    }
                    if (this.ws != null) this.ws.send(JSON.stringify({ user_id : this.user.id, name: HOME_EVENT_SONG_UPDATE,data:{}}))
                })
                .catch(errors => {
                    console.log(errors)
                }) 
        },
        sendPlayMusic(id){
            if (this.ws != null) this.ws.send(JSON.stringify({ user_id : this.user.id, name: HOME_EVENT_MUSIC_PLAY,data:{id : id}}))
        },
        sendStopMusic(){
            if (this.ws != null) this.ws.send(JSON.stringify({ user_id : this.user.id, name: HOME_EVENT_MUSIC_STOP,data:{}}))
        },
        sendChangeVolume(){
            if (this.ws != null) this.ws.send(JSON.stringify({ user_id : this.user.id, name: HOME_EVENT_VOLUME_CHANGE,data:{volume : this.media.volume }}))
        },
        updateDevice(device){
 
            axios
                .put(this.baseUrl() + '/api/v1/devices/' + device.id , device)
                .then(response => {
                    if (response.data.error != null){
                        return
                    }
                    if (this.ws != null) this.ws.send(JSON.stringify({ name: HOME_EVENT_DEVICE_UPDATE,data:{}}))

                })
                .catch(errors => {
                    console.log(errors)
                }) 
        },
        logout(){
            if (window.localStorage.getItem('session')) {
                window.localStorage.removeItem('session')
                window.location = this.baseUrl() + "/index.html" 
            }
        },
        loadSession(){
            if (!window.localStorage.getItem('session')) {
                window.location = this.baseUrl() + "/index.html"
                return;
            }

            this.user = JSON.parse(window.localStorage.getItem('session'))

            this.initAudioPlayer()
            this.loadDevices()
            this.loadMusics()
            this.initHomeWebsocket()

        },
        initHomeWebsocket(){
            this.ws = new WebSocket(this.baseWsUrl() + "/ws-home/" + this.makeid(15))
            this.ws.onmessage = (evt) => {
                let event = JSON.parse(evt.data)
                if (event.user_id != this.user.id){
                    return
                }

                switch (event.name) {
                    case  HOME_EVENT_DEVICE_UPDATE:

                        this.device.query.offset = 0
                        this.loadDevices()
                        break;
                    case HOME_EVENT_SONG_UPDATE:

                        this.music.query.offset = 0
                        this.loadMusics()
                        break;
                    case HOME_EVENT_DEVICE_LOGIN:
                        break;
                    case HOME_EVENT_DEVICE_LOGOUT:
                        break;
                    case HOME_EVENT_MUSIC_PLAY:

                        this.playAudioPlayer(event.data.id)
                        break;
                    case HOME_EVENT_VOLUME_CHANGE:

                        this.updateAudioPlayerVolume(event.data.volume)
                        break;       
                    case HOME_EVENT_MUSIC_STOP:

                        this.stopAudioPlayer()
                        break;
                    default: break;
                }
            }
            this.ws.onopen = () => {
                console.log("websoket open")   
            }
            this.ws.onclose = () => {
                console.log("websoket close")  
            }
            this.ws.onerror = (e) => {
                console.log(e)

            }
        },
        initAudioPlayer(){
            this.media.audio_player = this.$refs["audio-player"]

            let vm = this
            this.media.audio_player.onended = function() {

                axios
                    .get(vm.baseUrl() + '/api/v1/musics-random/' + vm.user.id)
                    .then(response => {
                        if (response.data.error != null){
                            return
                        }
                        if (response.data.data == null){
                            return
                        }

                        vm.sendPlayMusic(response.data.data.id)
                    })
                    .catch(errors => {
                        console.log(errors)
                    })

            }
            this.media.audio_player.volume = this.media.volume / 10
        },
        playAudioPlayer(song_id){

            axios
            .get(this.baseUrl() + '/api/v1/musics/' + song_id)
            .then(response => {
                if (response.data.error != null){
                    return
                }
                if (response.data.data == null){
                    return
                }

                this.media.current_song = response.data.data


                if (this.media.current_device.role == 0) {
                    return;
                }
                this.media.audio_player.load()
                this.media.audio_player.play()
            })
            .catch(errors => {
                console.log(errors)
            })


        },
        stopAudioPlayer(){
            this.media.audio_player.pause()
        },
        updateAudioPlayerVolume(volume){

            if (this.media.current_device.role == 0) {
                return;
            }
            this.media.volume = volume
            this.media.audio_player.volume = this.media.volume / 10
        },
        backPress(){
            if (event.state && event.state.noBackExitsApp) {
                window.history.pushState({ noBackExitsApp: true }, '')
            }
        },
        makeid(length) {
            let result           = '';
            let characters       = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
            let charactersLength = characters.length;
            for ( var i = 0; i < length; i++ ) {
                result += characters.charAt(Math.floor(Math.random() * charactersLength));
            }
            return result;
        },
        setCurrentHost(){
            this.host.name = window.location.hostname
            this.host.port = location.port
            this.host.protocol = location.protocol.concat("//")
            this.host.ws_protocol = this.host.protocol == "https://" ? "wss://"  : "ws://" 
        },
        baseUrl(){
            return this.host.protocol.concat(this.host.name + ":" + this.host.port)
        },
        baseWsUrl(){
            return this.host.ws_protocol.concat(this.host.name + ":" + this.host.port)
        }
    }
})