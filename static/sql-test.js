$(function() {

    var url = "/" + dbId + "/exec"
    
    function validate(line) {
        return line !== "";
    }

    function updateQuestion(next) {
        var $question = $("#question"),
            orig = $question.css("backgroundColor");
        $question.css({backgroundColor: "#E1F5C4"});
        $("#question .text").html("<span class='correct'>Correct!</span>");
        setTimeout(function() {
            $("#question .number").text(next.number);
            $("#question .text").text(next.text);
            $question.css({backgroundColor: orig});
        }, 3000);
    }

    function handle(line, report) {
        $.post(url, {
            id: dbId,
            sql: line
        }).done(function(data) {
            if (data.answered) {
                updateQuestion(data.next);
            }
            if (data.error) {
                report([{msg: data.error, className: "error"}]);
            } else {
                report([{msg: data.response, className: "success result"}]);
            }
        }).fail(function() {
            report([{msg: "An unknown error occured!", className: "error"}]);
        });
    }

    sqlConsole = $("#console").console({
        promptLabel: 'sqlite> ',
        commandValidate: validate,
        commandHandle: handle,
        autofocus: true,
        animateScroll: true,
        promptHistory: true,
        historyPreserveColumn: true
    });
});
