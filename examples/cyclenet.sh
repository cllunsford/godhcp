#!/bin/bash
while :
do
	ifdown eth1
	ifup eth1
done
