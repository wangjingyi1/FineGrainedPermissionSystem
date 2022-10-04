import axios from "axios";

export default {
    Get(api, cb) {
        axios.get(api)
            .then(cb)
            .catch(err => {
                console.log(err);
                this.$message(err)
            });
    },
    Post(api, post, cb) {
        axios.post(api, post)
            .then(cb)
            .catch(err => {
                console.log(err);
                this.$message(err)
            });
    },
};
