;(function () {
    let R = {
        options: {},
        heartElem: '<span class="reaction-btn__heart"/>',
        countElem: '<span class="reaction-btn__count"/>',
        countFormatter: (count) => {
            return (count >= R.options.minLikeToShowCount) ? count: '';
        },
        init: (options) => {
            options.apiURL = options.apiURL || 'http://reactions.mesuutt.com/api/v1';
            options.viewOnly = !!options.viewOnly;
            options.likedBtnClass = options.likedBtnClass? options.likedBtnClass : 'liked';

            options.loadCSS = options.loadCSS == undefined ? true: !!options.loadCSS;
            options.renderButton = options.renderButton? options.renderButton : R.renderButton;
            if (options.countFormatter) {
                R.countFormatter = options.countFormatter;
            }
            if (options.onLike) R.onLike = options.onLike;
            if (options.onUndoLike) R.onUndoLike = options.onUndoLike;

            R.options = options;

            if (R.options.loadCSS) {
                R.loadCSS("http://claps.test:9000/assets/like/like.css").then(function(){
                    if (R.options.loadUmbrealla){
                        R.loadJS("https://cdnjs.cloudflare.com/ajax/libs/umbrella/3.1.0/umbrella.min.js").then(R.run)
                    } else {
                        R.run();
                    }
                });
            } else {
                if (R.options.loadUmbrealla){
                    R.loadJS("https://cdnjs.cloudflare.com/ajax/libs/umbrella/3.1.0/umbrella.min.js").then(R.run)
                } else {
                    R.run();
                }
            }
        },
        loadJS: function(src, callback) {
            return new Promise(function( resolve, reject ) {
                var script = document.createElement('script');
                script.setAttribute('src', src);
                  script.onreadystatechange = script.onload = resolve;
                document.getElementsByTagName('head')[0].appendChild(script);
            });

        },
        loadCSS: function(src) {
            return new Promise(function( resolve, reject ) {
                var link = document.createElement( 'link' );
                link.rel  = 'stylesheet';
                link.href = src;
                document.head.appendChild( link );
                link.onload = resolve;
            });
        },
        run: function() {
            R.initializeFoundElems();
        },
        initializeFoundElems: function(){
            u(R.options.selector).each(function(item, i){
                item = u(item);
                R.api.getCount(item.attr('data-identifier'), item.attr('data-url')).then(
                    response => response.json()
                ).then(
                    data => R.renderButton(data, item)
                )
            });
        },
        onReactionChange: function(){
            var btn = this;
            var ident = u(btn).attr('data-identifier');

            var hasLikedBefore = localStorage.getItem(ident);
            if (hasLikedBefore) {
                R.api.decreaseCount(ident).then(
                    response => response.json()
                ).then(
                    data => {
                        R.removeLikeFromLocalStorage(ident);
                        if (data.error_message) {
                            console.log(data.error_message);
                            return;
                        }
                        R.onUndoLike(btn, data.count)
                    }
                )
            } else {
                R.api.increaseCount(ident).then(
                    response => response.json()
                ).then(
                    data => {
                        R.saveLikeToLocalStorage(ident);
                        if (data.error_message) {
                            console.log(data.error_message);
                            return;
                        }

                        R.onLike(btn, data.count);
                    }
                )
            }
        },
        onLike: function(btn, count) {
            u(btn).empty().append(u(R.heartElem));
            u(btn).empty().append(u(R.countElem).text(R.countFormatter(count)));
            u(btn).addClass('liked');
        },
        onUndoLike: function(btn, count) {
            u(btn).empty().append(u(R.heartElem));
            u(btn).empty().append(u(R.countElem).text(R.countFormatter(count)));
            u(btn).removeClass(R.options.likedBtnClass);
        },
        renderButton: function(data, btn) {
            u(btn).append(u(R.heartElem));
            u(btn).append(u(R.countElem).text(R.countFormatter(data.count)));
            if (R.options.viewOnly) {
                u(btn).first().style.display = '';
                return;
            };

            u(btn).on('click', R.onReactionChange);
            var ident = u(btn).attr('data-identifier');

            if (localStorage.getItem(ident)) {
                u(btn).addClass(R.options.likedBtnClass);
            }

            u(btn).first().style.display = '';
        },
        saveLikeToLocalStorage: function(uniqueIdent){
            localStorage.setItem(uniqueIdent, 'liked');
        },
        removeLikeFromLocalStorage: function(uniqueIdent){
            localStorage.removeItem(uniqueIdent);
        },
        api: {
            getCount: function(ident) {
                return fetch(R.options.apiURL + '/likes/count',{
                    method: 'POST',
                    headers: {'Content-Type': 'application/json',},
                    body: JSON.stringify({
                        'identifier': ident,
                        'page_url': window.location.href,
                    })
                });
            },
            increaseCount: function(ident) {
                return fetch(R.options.apiURL + '/likes/increase',{
                    method: 'POST',
                    headers: {'Content-Type': 'application/json',},
                    body: JSON.stringify({
                        'identifier': ident,
                        'page_url': window.location.href,
                    })
                });
            },
            decreaseCount: function(ident) {
                return fetch(R.options.apiURL + '/likes/decrease',{
                    method: 'POST',
                    headers: {'Content-Type': 'application/json',},
                    body: JSON.stringify({
                        'identifier': ident,
                        'page_url': window.location.href,
                    })
                });
            },
        }
    }

    document.addEventListener('DOMContentLoaded', function(event) {
        if (window.initReactionsManually) return;

        R.init({
            autoInit: true,
            loadUmbrealla: true,
            selector: '.reaction-btn',
            minLikeToShowCount: 1,
            // loadCSS: true,
            countFormatter: function (count) {
                console.log('M111');
                if (count) return count + " likes";
            },

            /*
            renderButton: function(){},
            onLike: function(btn, count){
                console.log('Liked', btn, count);
            },
            onUndoLike: function(btn, count){
                console.log('Undo liked', btn, count);
            }*/

            // likeBtnClass: 'my-like-btn',
            // likedBtnClass: 'my-liked-btn',

        });
    });

    window.Reactionist = R;
})();
