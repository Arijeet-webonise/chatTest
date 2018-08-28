const gulp = require('gulp'),
    sass = require('gulp-sass'),
    sourcemaps = require('gulp-sourcemaps'),
    autoprefixer = require('gulp-autoprefixer'),
    concat = require('gulp-concat'),
    symlink = require('gulp-symlink'),
    livereload = require('gulp-livereload'),
    babel = require('gulp-babel'),
    browserSync = require('browser-sync').create();

const path = {
    css : {
        src : './src/css/**/*.{sass,scss}',
        dest : './web/assets/css'
    },
    image : {
        src : './src/image',
        dest : './web/assets/image',
    },
    js : {
        src : './src/js/**/*.js',
        dest : './web/assets/js',
    },
    html : {
        src : './web/views/**/*.html',
    },
};

gulp.task('browser-sync', function(){
    browserSync.init({
        proxy: "http://localhost:9999"
    });
});

gulp.task('image', function(){
    return gulp.src(path.image.src)
        .pipe(symlink(path.image.dest));
});

gulp.task('js', function(){
    return gulp.src(path.js.src)
        .pipe(pipe(sourcemaps.init()))
        .pipe(concat('main.js'))
        .pipe(babel({
            presets: ['@babel/env']
        }))
        .pipe(sourcemaps.write("./maps"))
        .pipe(gulp.dest(path.js.dest))
        .pipe(livereload())
        .pipe(browserSync.stream());
});

gulp.task('html', function(){
    return gulp.src(path.html.src)
        .pipe(livereload())
        .pipe(browserSync.stream());
});

gulp.task('css', function(){
    return gulp.src(path.css.src)
        .pipe(sourcemaps.init())
        .pipe(sass.sync({outputStyle: 'compressed'}).on('error',sass.logError))
        .pipe(autoprefixer({
            browsers: ['last 2 version'],
            cascade: false
        }))
        .pipe(concat('style.css'))
        .pipe(sourcemaps.write("./maps"))
        .pipe(gulp.dest(path.css.dest))
        .pipe(livereload())
        .pipe(browserSync.stream());
});

gulp.task('watch', ['browser-sync','css', 'html'], function(){
    gulp.watch(path.css.src, ['css']);
    gulp.watch(path.html.src, ['html']);
});

gulp.task('default', ['css', 'image', 'node_modules']);
