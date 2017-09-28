/*
Template Name: Liebe
Author: Ingrid Kuhn
Author URI: themeforest/user/ingridk
Version: 1.0
*/

"use strict";
$(document).ready(function() {


    $("#story-carousel").owlCarousel({
        dots: true,
		 margin: 50,
        loop:true,
        autoplay: false,
        nav: true,
		  navText: [
            "<i class='fa fa-chevron-left'></i>",
            "<i class='fa fa-chevron-right'></i>"

        ],
        responsive: {
            1: {
                items: 1,
            },
			600: {
                items: 2,
            },
            1000: {
                items: 3,
            },
        }
    });


}); // end document ready

