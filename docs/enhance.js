function enhance() {
    // if there is no hash, start with the first slide
    var hash = window.location.hash;
    if (hash == "") {
	window.location += "#1"
	return
    }

    // handle clicks on slide changes
    window.onhashchange = toggleSlide;

    hash = hash.substring(1) // ignore the #
    var slides = document.getElementsByClassName("slide");
    for (var i = 0; i < slides.length; i++) {
	var s = slides.item(i);

	// hide all but the hashed one	
	if (s.id != hash) {
	    s.style.display = 'none';
	} else {
	    s.style.diplay = 'block';
	}
    }
}

var currentHash = window.location.hash;

function toggleSlide() {
    var c = document.getElementById(currentHash.substring(1));
    c.style.display = 'none';
    var n = document.getElementById(window.location.hash.substring(1))
    n.style.display = 'block';
    currentHash = window.location.hash;
}



enhance()
