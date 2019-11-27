$(document).ready(function() {

    // Check for click events on the navbar burger icon
    $("div.bingo-box").click(function(event) {
        console.log("bingo clicked");
        $target = $(event.target);
        $target.toggleClass('has-background-danger');
        $target.toggleClass('has-text-white');
    });
});