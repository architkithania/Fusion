# Fusion - A minimalist SDL2 Wrapper for Go

Fusion is a **VERY** minimalistic wrapper around SDL2 to wrap tedious functionalities of SDL2 using a convenient API. Note that Fusion does not aim to replace SDL2, rather it complements SDL2 by providing wrappers that help increase developer productivity. 



This project uses [this](https://github.com/veandco/go-sdl2) awesome project that provides a Go wrapper for the SDL2 library. Go check them out!

## Installation

Follow these instructions to setup the library for development. The project requires a few dependencies before since we will be working with Android as well. 

Please follow the steps [here](https://github.com/veandco/go-sdl2/blob/master/README.md). To set up the go-sdl2 as that is a direct dependency of this project.

You must also create a standalone android-ndk for `armeabi-v7a` and place it in the `creating_apk` folder with the folder name `toolchain`.



## Required Environment Variables

The following environment variables must be defined. 

```bash
export GOPATH=/path/to/go

export ANDROID_NDK_HOME=/path/to/android-ndk
export ANDROID_SYSROOT=${ANDROID_NDK_HOME}/platforms/android-16/arch-arm

export ANDROID_HOME=/path/to/android-sdk
export ANDROID_SDK_ROOT=/path/to/android-sdk
```



Furthermore, the path variable needs to be updated to include the following directory. If you do not want to clutter your path, you may define the following two lines in the the `tools/envs` folder. The scripts in tools automatically source that file when they need to build.

```bash
export PATH=$PATH:${ANDROID_NDK_HOME}/toolchains/arm-linux-androideabi-4.9/prebuilt/linux-x86_64/bin
export PATH=$PATH:/path/to/project/creating_apk/toolchain/bin

```



## Explanation of the various files

Let us go over the different files and folder of this project

1. `assets` : This folder is a requirement and must exist in order to use this library. The assets folder must contain folders with the names `fonts`, `images`, `music`, `sounds` (These folders may be empty but they must exist). All used fonts, images, music, and sounds should go in their respective folders
2. `creating_apk`: This folder contains everything to do with the android development system including the  `gradle` files as well as the `Java` source code. Generally you never need to interact with this folder except of when you need to get the built `apk` file for android development which gets built into the `creating_apk/android/android`

3. `src`: This is the folder where all the source code for this library exists

4. `tools`: This folder contains a set of folders that are useful for building for android. For building for android, run the tools in the following order: 

   1. `buildlib.sh`
   2. `buildapk.sh`
   3. `transferapk.sh`

   Please read the documentation in the tool scripts for more details.