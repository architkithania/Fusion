# This script is used to build the Java APK file. If successful the built APK can be found under
# creating_apk/android/android/build/outputs/apk/debug/*.apk
rm -rf android/src/main/assets

cp -rf ../../assets android/src/main/assets

./gradlew assembleDebug
