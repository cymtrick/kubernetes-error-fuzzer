#!/usr/bin/env perl
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
#
# Generate system call table for DragonFly from master list
# (for example, /usr/src/sys/kern/syscalls.master).

use strict;

if($ENV{'GOARCH'} eq "" || $ENV{'GOOS'} eq "") {
	print STDERR "GOARCH or GOOS not defined in environment\n";
	exit 1;
}

my $command = "mksysnum_dragonfly.pl " . join(' ', @ARGV);

print <<EOF;
// $command
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build $ENV{'GOARCH'},$ENV{'GOOS'}

package unix

const (
EOF

while(<>){
	if(/^([0-9]+)\s+STD\s+({ \S+\s+(\w+).*)$/){
		my $num = $1;
		my $proto = $2;
		my $name = "SYS_$3";
		$name =~ y/a-z/A-Z/;

		# There are multiple entries for enosys and nosys, so comment them out.
		if($name =~ /^SYS_E?NOSYS$/){
			$name = "// $name";
		}
		if($name eq 'SYS_SYS_EXIT'){
			$name = 'SYS_EXIT';
		}

		print "	$name = $num;  // $proto\n";
	}
}

print <<EOF;
)
EOF
