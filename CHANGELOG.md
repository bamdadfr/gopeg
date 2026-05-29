## [1.12.2](https://github.com/bamdadfr/gopeg/compare/v1.12.1...v1.12.2) (2026-05-29)


### Bug Fixes

* replace upscale x5 x6 with x2 x3 ([f41ab93](https://github.com/bamdadfr/gopeg/commit/f41ab93d513a4ab1e8c1cbe6e8577fa01c805b59))

## [1.12.1](https://github.com/bamdadfr/gopeg/compare/v1.12.0...v1.12.1) (2026-05-29)


### Bug Fixes

* actually add crf1 ([d3ff419](https://github.com/bamdadfr/gopeg/commit/d3ff419a1337ea1dd17a499cbd5c2a1cb10c6ff7))

# [1.12.0](https://github.com/bamdadfr/gopeg/compare/v1.11.0...v1.12.0) (2026-05-29)


### Features

* add 10 bits x264 and crf 1 ([f439f04](https://github.com/bamdadfr/gopeg/commit/f439f0401bce1364c32a0c58827b95965052f54d))


### Performance Improvements

* replace placebo x264 to veryslow ([15850ed](https://github.com/bamdadfr/gopeg/commit/15850edb6bf231668a4ddda2f497cafb1238f044))

# [1.11.0](https://github.com/bamdadfr/gopeg/compare/v1.10.0...v1.11.0) (2026-05-29)


### Bug Fixes

* rewrite insta v1 desc ([2afc2d1](https://github.com/bamdadfr/gopeg/commit/2afc2d1cd97f24f662d79728631e9c3de599e2af))


### Features

* add new presets ([1ff9f38](https://github.com/bamdadfr/gopeg/commit/1ff9f38143a2a1b14af9b2c7e50d2d6e65359f9e))

# [1.10.0](https://github.com/bamdadfr/gopeg/compare/v1.9.0...v1.10.0) (2026-05-29)


### Bug Fixes

* renames ([69efc30](https://github.com/bamdadfr/gopeg/commit/69efc3020414a3195690704ac10b4cedb8367ad9))


### Features

* add new archive declinations ([f6c9264](https://github.com/bamdadfr/gopeg/commit/f6c92644d86c00d5954e6d8cf477d87b934a9bed))
* add new video2x presets ([2a27d9a](https://github.com/bamdadfr/gopeg/commit/2a27d9a2f271d2db8488afee731ac4c1a73f6270))
* rename instagram presets and add square v2 ([f25c7cc](https://github.com/bamdadfr/gopeg/commit/f25c7cc441b6e01af1f441813a11e77dd55d1ae9))

# [1.9.0](https://github.com/bamdadfr/gopeg/compare/v1.8.0...v1.9.0) (2026-05-27)


### Bug Fixes

* add better description ([13f7f18](https://github.com/bamdadfr/gopeg/commit/13f7f18d0dae9babb5aaa29f48e471cfb23c67f2))
* derive output paths from fields ([70ba1fc](https://github.com/bamdadfr/gopeg/commit/70ba1fc21de5ec81929e546d94b5c9a9c4e86a69)), closes [#10](https://github.com/bamdadfr/gopeg/issues/10)


### Features

* add new presets ([d046f4a](https://github.com/bamdadfr/gopeg/commit/d046f4a4a2649631eb0784699d8a35f947c8e860))

# [1.8.0](https://github.com/bamdadfr/gopeg/compare/v1.7.1...v1.8.0) (2026-05-26)


### Bug Fixes

* track context to kill all routines on quit ([0adb0d4](https://github.com/bamdadfr/gopeg/commit/0adb0d4824521214624d95c764040a3a28ce97de))


### Features

* UI subscription to queue ([a3bfc1b](https://github.com/bamdadfr/gopeg/commit/a3bfc1bf8bdb0958e3169f337db64d0daab8e8c9))

## [1.7.1](https://github.com/bamdadfr/gopeg/compare/v1.7.0...v1.7.1) (2026-05-26)


### Bug Fixes

* add mutext to queue ([b1cd97c](https://github.com/bamdadfr/gopeg/commit/b1cd97c80b5631d032d4808c213a51cd5f2ad22d))

# [1.7.0](https://github.com/bamdadfr/gopeg/compare/v1.6.0...v1.7.0) (2026-05-26)


### Features

* add draft UI for queue ([146e6ed](https://github.com/bamdadfr/gopeg/commit/146e6ed185718e97329931a4da57af449369e80a))

# [1.6.0](https://github.com/bamdadfr/gopeg/compare/v1.5.5...v1.6.0) (2026-05-26)


### Bug Fixes

* relocate binary validation to its package ([ce6ca04](https://github.com/bamdadfr/gopeg/commit/ce6ca042fb1bb9d2fcac5575ae0fa932bada8745))


### Features

* relocate and refactor ([d7a57b7](https://github.com/bamdadfr/gopeg/commit/d7a57b705f3653b9ea4120754a96f61a1d87e4a0))

## [1.5.5](https://github.com/bamdadfr/gopeg/compare/v1.5.4...v1.5.5) (2026-05-25)


### Performance Improvements

* extract binary to its own struct ([f833a76](https://github.com/bamdadfr/gopeg/commit/f833a7666a34816a552cdac44093bb4629f6e977))
* refactor command and args ([8ef0909](https://github.com/bamdadfr/gopeg/commit/8ef09090a5b332b0693899ee2d14be2937126e9d))
* relocate preset list creation ([6ecea49](https://github.com/bamdadfr/gopeg/commit/6ecea496f608c4f2a5ad099edc4d7db3a4b6ee3b))

## [1.5.4](https://github.com/bamdadfr/gopeg/compare/v1.5.3...v1.5.4) (2026-05-25)


### Bug Fixes

* replace text objects with labels and use theme to apply font size ([3eb5455](https://github.com/bamdadfr/gopeg/commit/3eb5455a7c20ae38cbdc8e4bc821d85b2956ed86))
* use new sorting native packages ([d72f7c5](https://github.com/bamdadfr/gopeg/commit/d72f7c583df2b9439ae2e030a63517de495b4ccf))

## [1.5.3](https://github.com/bamdadfr/gopeg/compare/v1.5.2...v1.5.3) (2026-05-25)


### Bug Fixes

* improve font size ([63e9e37](https://github.com/bamdadfr/gopeg/commit/63e9e37e8bbb28637f34c65cb93af5656788348f))
* remove UI refresh from go routine ([c9a6c6b](https://github.com/bamdadfr/gopeg/commit/c9a6c6bd054e7cf0681fdf29c55a1dc3a45c416f))

## [1.5.2](https://github.com/bamdadfr/gopeg/compare/v1.5.1...v1.5.2) (2026-05-25)


### Bug Fixes

* display cmd error in status ([a4fe789](https://github.com/bamdadfr/gopeg/commit/a4fe789b9bde55caabe40824e2815ea98919c1a3))
* extract binary name getter ([44401d4](https://github.com/bamdadfr/gopeg/commit/44401d451c83b483ab39f510b7474fc184e9421a))
* typo ([8ddce35](https://github.com/bamdadfr/gopeg/commit/8ddce358305508d87a15944200c6f742cd4f3093))
* use atomic isRunning env var ([3e962fd](https://github.com/bamdadfr/gopeg/commit/3e962fd6acfb3f633376e41098049878c0106a55))


### Performance Improvements

* split main package ([af3b0e7](https://github.com/bamdadfr/gopeg/commit/af3b0e7108ba6cc9b59d296ea82054760cbc49ea))

## [1.5.1](https://github.com/bamdadfr/gopeg/compare/v1.5.0...v1.5.1) (2026-05-25)


### Performance Improvements

* remove dead code ([033530b](https://github.com/bamdadfr/gopeg/commit/033530b342acac20e666787917dd50e91be513af))

# [1.5.0](https://github.com/bamdadfr/gopeg/compare/v1.4.0...v1.5.0) (2026-05-25)


### Bug Fixes

* only append -y for existing output for ffmpeg commands ([4ef6f7d](https://github.com/bamdadfr/gopeg/commit/4ef6f7d88e4c05aaae8ba72f286e5b3c02d73365))
* watch for already running and run into goroutine ([35f09cb](https://github.com/bamdadfr/gopeg/commit/35f09cb9aa8477002e038ab7f9e25a0b8b817852))


### Features

* windows path for video2x ([dae70de](https://github.com/bamdadfr/gopeg/commit/dae70dee9404dd661b3f49d2f1616b40dc2653be))

# [1.4.0](https://github.com/bamdadfr/gopeg/compare/v1.3.0...v1.4.0) (2026-05-25)


### Features

* add RIFE 4x preset ([a63d885](https://github.com/bamdadfr/gopeg/commit/a63d88567aa0b9a7e228f9e82db1cddde6b63faf))


### Performance Improvements

* extract status setter ([fb160d7](https://github.com/bamdadfr/gopeg/commit/fb160d73db8866c856f29ae49a34755c303e81c1))

# [1.3.0](https://github.com/bamdadfr/gopeg/compare/v1.2.0...v1.3.0) (2026-05-25)


### Features

* add status bar and bin detection ([7d7a3ab](https://github.com/bamdadfr/gopeg/commit/7d7a3abc5771bf3649247f52450676194762dbca))

# [1.2.0](https://github.com/bamdadfr/gopeg/compare/v1.1.1...v1.2.0) (2026-05-25)


### Features

* add new presets and refactor ([fead695](https://github.com/bamdadfr/gopeg/commit/fead6953a938c4fbb59c0068f44c09ae452aa681))

## [1.1.1](https://github.com/bamdadfr/gopeg/compare/v1.1.0...v1.1.1) (2026-05-23)


### Bug Fixes

* typo ([ab14585](https://github.com/bamdadfr/gopeg/commit/ab1458518e553169e59dfd00736158f214fa63bd))

# [1.1.0](https://github.com/bamdadfr/gopeg/compare/v1.0.0...v1.1.0) (2026-05-23)


### Features

* add archive preset ([e08a332](https://github.com/bamdadfr/gopeg/commit/e08a332a50df4b057238c21fee289288c638aae6))
* add preset sorting ([69d555b](https://github.com/bamdadfr/gopeg/commit/69d555b1399aa4f9a8c6ca96f61608af76a52493))

# 1.0.0 (2026-05-22)


### Features

* init ([1d56e67](https://github.com/bamdadfr/gopeg/commit/1d56e673a1edb2311239cef860f13143d19e94d2))
