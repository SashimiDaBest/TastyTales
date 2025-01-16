import { greet } from './utils.js';

document.getElementById("myButton").addEventListener("click", () => {
    greet("User");
});

document.addEventListener("DOMContentLoaded", () => {
    const button = document.getElementById("myButton");
    if (button) {
        button.addEventListener("click", () => {
            console.log("Button clicked! JS working.");
        });
    } else {
        console.error("Button not found.");
    }
});