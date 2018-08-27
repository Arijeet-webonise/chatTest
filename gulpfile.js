const gulp = require('gulp'),
    sass = require('gulp-sass'),
    sourcemaps = require('gulp-sourcemaps'),
    autoprefixer = require('gulp-autoprefixer'),
    concat = require('gulp-concat');

const path = {
    css : {
        src : './src/css/**/*.{sass,scss}',
        dest : './web/assets/css'
    },
};


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
        .pipe(gulp.dest(path.css.dest));
});

gulp.task('watch', ['css'], function(){
    gulp.watch(path.css.src, ['css']);
});

gulp.task('default', ['css']);
