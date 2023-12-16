// address: https://api.github.com/users/amirhnajafiz-learning/repos
// helper func: https://stackoverflow.com/questions/35442329/visualizing-readme-md-files-in-my-website

async function repositories() {
    return (username) => {
        fetch(`https://api.github.com/users/${username}/repos`)
            .then((response) => response.json())
            .then((data) => {
                list = [];

                data.forEach(el => {
                    list.push({
                        "id": el['id'],
                        "name": el['name'],
                        "description": el['description']
                    });
                });

                return list;
            })
            .catch((e) => {
                console.log(e);

                return [];
            })
    }
}

async function readme() {

}

export const api = {
    pull: repositories(),
    read: readme()
}
