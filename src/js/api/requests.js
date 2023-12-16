// address: https://api.github.com/users/amirhnajafiz-learning/repos
// sub address: https://raw.githubusercontent.com/amirhnajafiz-learning/ansible/main/README.md
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
                        "description": el['description'],
                        "branch": el['default_branch'],
                        "topics": el['topics']
                    });
                });

                return list;
            })
            .catch((e) => {
                console.error(e);

                return [];
            })
    }
}

async function readme() {
    return (username, repository, branch) => {
        fetch(`https://raw.githubusercontent.com/${username}/${repository}/${branch}/README.md`)
            .then((response) => response.text())
            .then((date) => date)
            .catch((e) => {
                console.error(e);

                return "";
            })
    }
}

export const api = {
    pull: repositories(),
    read: readme()
}
