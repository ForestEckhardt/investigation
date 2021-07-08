## Findings

Let's talk about the general findings of this exercise. For the sample that I
used which was the latest versions of the following dependencies:
* bundler
* curl
* dotnet-aspnetcore
* dotnet-runtime
* dotnet-sdk
* go
* httpd
* icu
* nginx
* node
* php
* pip
* pipenv
* python
* ruby
* rust
* tini
* yarn

Only a hand full of these had strange extraneous licenses that I had trouble
locating in the source (`curl` having `SSH-Short` is an example of this kind of
inconsistency). None of these dependencies were missing their actual license
form the list of licenses at all. My recommendation for using this license
information is for our process to include all licenses detected and when doing
the compliance verification process violating licenses are verified by hand to
ensure that they are present in our dependency. I could not find what I felt
was any reasonable confidence cutoff score without being worried I would miss
license information that was present but the scanner was not confident in.

This should still have a have a net positive effect on the speed of compliance
as anything that is passing should be able to be waved through as at most the
dependency appears to only have extra licenses and the alternative is to hand
verify the license information which appears to be the existing process.
