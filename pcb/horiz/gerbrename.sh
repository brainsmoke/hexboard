#!/bin/sh
unset LC_TELEPHONE
rename s'/(.*?)-Edge.Cuts\..*/$1.GKO/' gerb/*
rename s'/(.*?)-.*.g2/$1.G2L/' gerb/*
rename s'/(.*?)-.*.g3/$1.G3L/' gerb/*
rename s'/(.*?)-.*\.(...)$/$1.\U$2/' gerb/*
rename s'/(.*?).drl/$1.XLN/' gerb/*
rename s'/hexdump/column/' gerb/*

