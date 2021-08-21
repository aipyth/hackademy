const { watch, series, src, dest } = require('gulp');

var browserSync = require('browser-sync').create();
var postcss = require('gulp-postcss');

const source = './src/'
const destination = './assets/'

function css(cb) {
    return src(source + '*.css')
        .pipe(postcss())
        .pipe(dest(destination + 'css'))
        .pipe(browserSync.stream());
    cb();
}

function html(cb) {
    return src(source + '*.html')
        .pipe(dest(destination));
    cb();
}

function browserSyncServe(cb) {
    browserSync.init({
        server: {
            baseDir: destination,
        },
    });
    cb();
}

function browserSyncReload(cb) {
    browserSync.reload();
    cb();
}

function watchTask() {
    watch(source + '*.html', series(html, browserSyncReload));
    watch([source + '*.css'], series(css, browserSyncReload));
}

exports.default = series(html, css, browserSyncServe, watchTask);
exports.css = css;
exports.build = series(html, css);
