// load plugins
import editor from "./plugins/markdown.js";
import { api } from "./api/requests.js";

// load components
import getSlideComponent from "./components/slide.js";

// variables
const user = "amirhnajafiz-learning";

// logic
async function main() {
    const repos = await api.pull(user);
    const wrapper = document.createElement("div");

    console.log(repos);

    repos.forEach(async repo => {
        let txt = await api.read(user, repo['name'], repo['branch']);
        let slide = editor(getSlideComponent(), txt);

        wrapper.appendChild(slide);
        wrapper.appendChild(document.createElement("hr"));
    });

    // mount component to main app
    document.getElementById("app").appendChild(wrapper);
}

await main();