val kotlinVersion = "1.3.40"

plugins {
    application
    kotlin("jvm") version "1.3.40"
}

group = "doubleStringReverse"
version = "1.0-SNAPSHOT"

//val mainClass by extra("io.ktor.server.netty.EngineMain")

application {
//    mainClassName = mainClass
}

java.sourceCompatibility = JavaVersion.VERSION_11

dependencies {
    implementation(kotlin("stdlib"))
    testCompile("junit:junit:4.12")
    testCompile("org.jetbrains.kotlin:kotlin-test-junit:$kotlinVersion")
}

repositories {
    jcenter()
}
