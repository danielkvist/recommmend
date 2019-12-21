require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

$(() => {});
$("body").click(function(){  
    $(".alert").alert("close");
});