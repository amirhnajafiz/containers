// load plugins
import editor from "./plugins/markdown.js";
import { api } from "./api/requests.js";

// load components
import getSlideComponent from "./components/slide.js";

// variables
const user = "amirhnajafiz-learning";

// logic
const repos = await api.pull(user);
const wrapper = document.createElement("div");

repos.forEach(repo => {
    let txt = api.read(user, repo['name'], repo['branch']);
    let slide = editor(getSlideComponent(), txt);

    wrapper.appendChild(slide);
});

// mount component to main app
document.getElementById("app").innerHTML = wrapper;
