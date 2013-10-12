$(function() {

    var url = "/" + dbId + "/exec"
    
    function validate(line) {
        return line !== "";
    }

    function handle(line, report) {
        $.post(url, {
            id: dbId,
            sql: line
        }).done(function(data) {
            console.log(data);
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
