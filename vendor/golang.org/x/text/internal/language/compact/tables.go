// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package compact

import "golang.org/x/text/internal/language"

// CLDRVersion is the CLDR version from which the tables in this package are derived.
const CLDRVersion = "32"

// NumCompactTags is the number of common tags. The maximum tag is
// NumCompactTags-1.
const NumCompactTags = 775
const (
	undIndex          ID = 0
	afIndex           ID = 1
	afNAIndex         ID = 2
	afZAIndex         ID = 3
	agqIndex          ID = 4
	agqCMIndex        ID = 5
	akIndex           ID = 6
	akGHIndex         ID = 7
	amIndex           ID = 8
	amETIndex         ID = 9
	arIndex           ID = 10
	ar001Index        ID = 11
	arAEIndex         ID = 12
	arBHIndex         ID = 13
	arDJIndex         ID = 14
	arDZIndex         ID = 15
	arEGIndex         ID = 16
	arEHIndex         ID = 17
	arERIndex         ID = 18
	arILIndex         ID = 19
	arIQIndex         ID = 20
	arJOIndex         ID = 21
	arKMIndex         ID = 22
	arKWIndex         ID = 23
	arLBIndex         ID = 24
	arLYIndex         ID = 25
	arMAIndex         ID = 26
	arMRIndex         ID = 27
	arOMIndex         ID = 28
	arPSIndex         ID = 29
	arQAIndex         ID = 30
	arSAIndex         ID = 31
	arSDIndex         ID = 32
	arSOIndex         ID = 33
	arSSIndex         ID = 34
	arSYIndex         ID = 35
	arTDIndex         ID = 36
	arTNIndex         ID = 37
	arYEIndex         ID = 38
	arsIndex          ID = 39
	asIndex           ID = 40
	asINIndex         ID = 41
	asaIndex          ID = 42
	asaTZIndex        ID = 43
	astIndex          ID = 44
	astESIndex        ID = 45
	azIndex           ID = 46
	azCyrlIndex       ID = 47
	azCyrlAZIndex     ID = 48
	azLatnIndex       ID = 49
	azLatnAZIndex     ID = 50
	basIndex          ID = 51
	basCMIndex        ID = 52
	beIndex           ID = 53
	beBYIndex         ID = 54
	bemIndex          ID = 55
	bemZMIndex        ID = 56
	bezIndex          ID = 57
	bezTZIndex        ID = 58
	bgIndex           ID = 59
	bgBGIndex         ID = 60
	bhIndex           ID = 61
	bmIndex           ID = 62
	bmMLIndex         ID = 63
	bnIndex           ID = 64
	bnBDIndex         ID = 65
	bnINIndex         ID = 66
	boIndex           ID = 67
	boCNIndex         ID = 68
	boINIndex         ID = 69
	brIndex           ID = 70
	brFRIndex         ID = 71
	brxIndex          ID = 72
	brxINIndex        ID = 73
	bsIndex           ID = 74
	bsCyrlIndex       ID = 75
	bsCyrlBAIndex     ID = 76
	bsLatnIndex       ID = 77
	bsLatnBAIndex     ID = 78
	caIndex           ID = 79
	caADIndex         ID = 80
	caESIndex         ID = 81
	caFRIndex         ID = 82
	caITIndex         ID = 83
	ccpIndex          ID = 84
	ccpBDIndex        ID = 85
	ccpINIndex        ID = 86
	ceIndex           ID = 87
	ceRUIndex         ID = 88
	cggIndex          ID = 89
	cggUGIndex        ID = 90
	chrIndex          ID = 91
	chrUSIndex        ID = 92
	ckbIndex          ID = 93
	ckbIQIndex        ID = 94
	ckbIRIndex        ID = 95
	csIndex           ID = 96
	csCZIndex         ID = 97
	cuIndex           ID = 98
	cuRUIndex         ID = 99
	cyIndex           ID = 100
	cyGBIndex         ID = 101
	daIndex           ID = 102
	daDKIndex         ID = 103
	daGLIndex         ID = 104
	davIndex          ID = 105
	davKEIndex        ID = 106
	deIndex           ID = 107
	deATIndex         ID = 108
	deBEIndex         ID = 109
	deCHIndex         ID = 110
	deDEIndex         ID = 111
	deITIndex         ID = 112
	deLIIndex         ID = 113
	deLUIndex         ID = 114
	djeIndex          ID = 115
	djeNEIndex        ID = 116
	dsbIndex          ID = 117
	dsbDEIndex        ID = 118
	duaIndex          ID = 119
	duaCMIndex        ID = 120
	dvIndex           ID = 121
	dyoIndex          ID = 122
	dyoSNIndex        ID = 123
	dzIndex           ID = 124
	dzBTIndex         ID = 125
	ebuIndex          ID = 126
	ebuKEIndex        ID = 127
	eeIndex           ID = 128
	eeGHIndex         ID = 129
	eeTGIndex         ID = 130
	elIndex           ID = 131
	elCYIndex         ID = 132
	elGRIndex         ID = 133
	enIndex           ID = 134
	en001Index        ID = 135
	en150Index        ID = 136
	enAGIndex         ID = 137
	enAIIndex         ID = 138
	enASIndex         ID = 139
	enATIndex         ID = 140
	enAUIndex         ID = 141
	enBBIndex         ID = 142
	enBEIndex         ID = 143
	enBIIndex         ID = 144
	enBMIndex         ID = 145
	enBSIndex         ID = 146
	enBWIndex         ID = 147
	enBZIndex         ID = 148
	enCAIndex         ID = 149
	enCCIndex         ID = 150
	enCHIndex         ID = 151
	enCKIndex         ID = 152
	enCMIndex         ID = 153
	enCXIndex         ID = 154
	enCYIndex         ID = 155
	enDEIndex         ID = 156
	enDGIndex         ID = 157
	enDKIndex         ID = 158
	enDMIndex         ID = 159
	enERIndex         ID = 160
	enFIIndex         ID = 161
	enFJIndex         ID = 162
	enFKIndex         ID = 163
	enFMIndex         ID = 164
	enGBIndex         ID = 165
	enGDIndex         ID = 166
	enGGIndex         ID = 167
	enGHIndex         ID = 168
	enGIIndex         ID = 169
	enGMIndex         ID = 170
	enGUIndex         ID = 171
	enGYIndex         ID = 172
	enHKIndex         ID = 173
	enIEIndex         ID = 174
	enILIndex         ID = 175
	enIMIndex         ID = 176
	enINIndex         ID = 177
	enIOIndex         ID = 178
	enJEIndex         ID = 179
	enJMIndex         ID = 180
	enKEIndex         ID = 181
	enKIIndex         ID = 182
	enKNIndex         ID = 183
	enKYIndex         ID = 184
	enLCIndex         ID = 185
	enLRIndex         ID = 186
	enLSIndex         ID = 187
	enMGIndex         ID = 188
	enMHIndex         ID = 189
	enMOIndex         ID = 190
	enMPIndex         ID = 191
	enMSIndex         ID = 192
	enMTIndex         ID = 193
	enMUIndex         ID = 194
	enMWIndex         ID = 195
	enMYIndex         ID = 196
	enNAIndex         ID = 197
	enNFIndex         ID = 198
	enNGIndex         ID = 199
	enNLIndex         ID = 200
	enNRIndex         ID = 201
	enNUIndex         ID = 202
	enNZIndex         ID = 203
	enPGIndex         ID = 204
	enPHIndex         ID = 205
	enPKIndex         ID = 206
	enPNIndex         ID = 207
	enPRIndex         ID = 208
	enPWIndex         ID = 209
	enRWIndex         ID = 210
	enSBIndex         ID = 211
	enSCIndex         ID = 212
	enSDIndex         ID = 213
	enSEIndex         ID = 214
	enSGIndex         ID = 215
	enSHIndex         ID = 216
	enSIIndex         ID = 217
	enSLIndex         ID = 218
	enSSIndex         ID = 219
	enSXIndex         ID = 220
	enSZIndex         ID = 221
	enTCIndex         ID = 222
	enTKIndex         ID = 223
	enTOIndex         ID = 224
	enTTIndex         ID = 225
	enTVIndex         ID = 226
	enTZIndex         ID = 227
	enUGIndex         ID = 228
	enUMIndex         ID = 229
	enUSIndex         ID = 230
	enVCIndex         ID = 231
	enVGIndex         ID = 232
	enVIIndex         ID = 233
	enVUIndex         ID = 234
	enWSIndex         ID = 235
	enZAIndex         ID = 236
	enZMIndex         ID = 237
	enZWIndex         ID = 238
	eoIndex           ID = 239
	eo001Index        ID = 240
	esIndex           ID = 241
	es419Index        ID = 242
	esARIndex         ID = 243
	esBOIndex         ID = 244
	esBRIndex         ID = 245
	esBZIndex         ID = 246
	esCLIndex         ID = 247
	esCOIndex         ID = 248
	esCRIndex         ID = 249
	esCUIndex         ID = 250
	esDOIndex         ID = 251
	esEAIndex         ID = 252
	esECIndex         ID = 253
	esESIndex         ID = 254
	esGQIndex         ID = 255
	esGTIndex         ID = 256
	esHNIndex         ID = 257
	esICIndex         ID = 258
	esMXIndex         ID = 259
	esNIIndex         ID = 260
	esPAIndex         ID = 261
	esPEIndex         ID = 262
	esPHIndex         ID = 263
	esPRIndex         ID = 264
	esPYIndex         ID = 265
	esSVIndex         ID = 266
	esUSIndex         ID = 267
	esUYIndex         ID = 268
	esVEIndex         ID = 269
	etIndex           ID = 270
	etEEIndex         ID = 271
	euIndex           ID = 272
	euESIndex         ID = 273
	ewoIndex          ID = 274
	ewoCMIndex        ID = 275
	faIndex           ID = 276
	faAFIndex         ID = 277
	faIRIndex         ID = 278
	ffIndex           ID = 279
	ffCMIndex         ID = 280
	ffGNIndex         ID = 281
	ffMRIndex         ID = 282
	ffSNIndex         ID = 283
	fiIndex           ID = 284
	fiFIIndex         ID = 285
	filIndex          ID = 286
	filPHIndex        ID = 287
	foIndex           ID = 288
	foDKIndex         ID = 289
	foFOIndex         ID = 290
	frIndex           ID = 291
	frBEIndex         ID = 292
	frBFIndex         ID = 293
	frBIIndex         ID = 294
	frBJIndex         ID = 295
	frBLIndex         ID = 296
	frCAIndex         ID = 297
	frCDIndex         ID = 298
	frCFIndex         ID = 299
	frCGIndex         ID = 300
	frCHIndex         ID = 301
	frCIIndex         ID = 302
	frCMIndex         ID = 303
	frDJIndex         ID = 304
	frDZIndex         ID = 305
	frFRIndex         ID = 306
	frGAIndex         ID = 307
	frGFIndex         ID = 308
	frGNIndex         ID = 309
	frGPIndex         ID = 310
	frGQIndex         ID = 311
	frHTIndex         ID = 312
	frKMIndex         ID = 313
	frLUIndex         ID = 314
	frMAIndex         ID = 315
	frMCIndex         ID = 316
	frMFIndex         ID = 317
	frMGIndex         ID = 318
	frMLIndex         ID = 319
	frMQIndex         ID = 320
	frMRIndex         ID = 321
	frMUIndex         ID = 322
	frNCIndex         ID = 323
	frNEIndex         ID = 324
	frPFIndex         ID = 325
	frPMIndex         ID = 326
	frREIndex         ID = 327
	frRWIndex         ID = 328
	frSCIndex         ID = 329
	frSNIndex         ID = 330
	frSYIndex         ID = 331
	frTDIndex         ID = 332
	frTGIndex         ID = 333
	frTNIndex         ID = 334
	frVUIndex         ID = 335
	frWFIndex         ID = 336
	frYTIndex         ID = 337
	furIndex          ID = 338
	furITIndex        ID = 339
	fyIndex           ID = 340
	fyNLIndex         ID = 341
	gaIndex           ID = 342
	gaIEIndex         ID = 343
	gdIndex           ID = 344
	gdGBIndex         ID = 345
	glIndex           ID = 346
	glESIndex         ID = 347
	gswIndex          ID = 348
	gswCHIndex        ID = 349
	gswFRIndex        ID = 350
	gswLIIndex        ID = 351
	guIndex           ID = 352
	guINIndex         ID = 353
	guwIndex          ID = 354
	guzIndex          ID = 355
	guzKEIndex        ID = 356
	gvIndex           ID = 357
	gvIMIndex         ID = 358
	haIndex           ID = 359
	haGHIndex         ID = 360
	haNEIndex         ID = 361
	haNGIndex         ID = 362
	hawIndex          ID = 363
	hawUSIndex        ID = 364
	heIndex           ID = 365
	heILIndex         ID = 366
	hiIndex           ID = 367
	hiINIndex         ID = 368
	hrIndex           ID = 369
	hrBAIndex         ID = 370
	hrHRIndex         ID = 371
	hsbIndex          ID = 372
	hsbDEIndex        ID = 373
	huIndex           ID = 374
	huHUIndex         ID = 375
	hyIndex           ID = 376
	hyAMIndex         ID = 377
	idIndex           ID = 378
	idIDIndex         ID = 379
	igIndex           ID = 380
	igNGIndex         ID = 381
	iiIndex           ID = 382
	iiCNIndex         ID = 383
	inIndex           ID = 384
	ioIndex           ID = 385
	isIndex           ID = 386
	isISIndex         ID = 387
	itIndex           ID = 388
	itCHIndex         ID = 389
	itITIndex         ID = 390
	itSMIndex         ID = 391
	itVAIndex         ID = 392
	iuIndex           ID = 393
	iwIndex           ID = 394
	jaIndex           ID = 395
	jaJPIndex         ID = 396
	jboIndex          ID = 397
	jgoIndex          ID = 398
	jgoCMIndex        ID = 399
	jiIndex           ID = 400
	jmcIndex          ID = 401
	jmcTZIndex        ID = 402
	jvIndex           ID = 403
	jwIndex           ID = 404
	kaIndex           ID = 405
	kaGEIndex         ID = 406
	kabIndex          ID = 407
	kabDZIndex        ID = 408
	kajIndex          ID = 409
	kamIndex          ID = 410
	kamKEIndex        ID = 411
	kcgIndex          ID = 412
	kdeIndex          ID = 413
	kdeTZIndex        ID = 414
	keaIndex          ID = 415
	keaCVIndex        ID = 416
	khqIndex          ID = 417
	khqMLIndex        ID = 418
	kiIndex           ID = 419
	kiKEIndex         ID = 420
	kkIndex           ID = 421
	kkKZIndex         ID = 422
	kkjIndex          ID = 423
	kkjCMIndex        ID = 424
	klIndex           ID = 425
	klGLIndex         ID = 426
	klnIndex          ID = 427
	klnKEIndex        ID = 428
	kmIndex           ID = 429
	kmKHIndex         ID = 430
	knIndex           ID = 431
	knINIndex         ID = 432
	koIndex           ID = 433
	koKPIndex         ID = 434
	koKRIndex         ID = 435
	kokIndex          ID = 436
	kokINIndex        ID = 437
	ksIndex           ID = 438
	ksINIndex         ID = 439
	ksbIndex          ID = 440
	ksbTZIndex        ID = 441
	ksfIndex          ID = 442
	ksfCMIndex        ID = 443
	kshIndex          ID = 444
	kshDEIndex        ID = 445
	kuIndex           ID = 446
	kwIndex           ID = 447
	kwGBIndex         ID = 448
	kyIndex           ID = 449
	kyKGIndex         ID = 450
	lagIndex          ID = 451
	lagTZIndex        ID = 452
	lbIndex           ID = 453
	lbLUIndex         ID = 454
	lgIndex           ID = 455
	lgUGIndex         ID = 456
	lktIndex          ID = 457
	lktUSIndex        ID = 458
	lnIndex           ID = 459
	lnAOIndex         ID = 460
	lnCDIndex         ID = 461
	lnCFIndex         ID = 462
	lnCGIndex         ID = 463
	loIndex           ID = 464
	loLAIndex         ID = 465
	lrcIndex          ID = 466
	lrcIQIndex        ID = 467
	lrcIRIndex        ID = 468
	ltIndex           ID = 469
	ltLTIndex         ID = 470
	luIndex           ID = 471
	luCDIndex         ID = 472
	luoIndex          ID = 473
	luoKEIndex        ID = 474
	luyIndex          ID = 475
	luyKEIndex        ID = 476
	lvIndex           ID = 477
	lvLVIndex         ID = 478
	masIndex          ID = 479
	masKEIndex        ID = 480
	masTZIndex        ID = 481
	merIndex          ID = 482
	merKEIndex        ID = 483
	mfeIndex          ID = 484
	mfeMUIndex        ID = 485
	mgIndex           ID = 486
	mgMGIndex         ID = 487
	mghIndex          ID = 488
	mghMZIndex        ID = 489
	mgoIndex          ID = 490
	mgoCMIndex        ID = 491
	mkIndex           ID = 492
	mkMKIndex         ID = 493
	mlIndex           ID = 494
	mlINIndex         ID = 495
	mnIndex           ID = 496
	mnMNIndex         ID = 497
	moIndex           ID = 498
	mrIndex           ID = 499
	mrINIndex         ID = 500
	msIndex           ID = 501
	msBNIndex         ID = 502
	msMYIndex         ID = 503
	msSGIndex         ID = 504
	mtIndex           ID = 505
	mtMTIndex         ID = 506
	muaIndex          ID = 507
	muaCMIndex        ID = 508
	myIndex           ID = 509
	myMMIndex         ID = 510
	mznIndex          ID = 511
	mznIRIndex        ID = 512
	nahIndex          ID = 513
	naqIndex          ID = 514
	naqNAIndex        ID = 515
	nbIndex           ID = 516
	nbNOIndex         ID = 517
	nbSJIndex         ID = 518
	ndIndex           ID = 519
	ndZWIndex         ID = 520
	ndsIndex          ID = 521
	ndsDEIndex        ID = 522
	ndsNLIndex        ID = 523
	neIndex           ID = 524
	neINIndex         ID = 525
	neNPIndex         ID = 526
	nlIndex           ID = 527
	nlAWIndex         ID = 528
	nlBEIndex         ID = 529
	nlBQIndex         ID = 530
	nlCWIndex         ID = 531
	nlNLIndex         ID = 532
	nlSRIndex         ID = 533
	nlSXIndex         ID = 534
	nmgIndex          ID = 535
	nmgCMIndex        ID = 536
	nnIndex           ID = 537
	nnNOIndex         ID = 538
	nnhIndex          ID = 539
	nnhCMIndex        ID = 540
	noIndex           ID = 541
	nqoIndex          ID = 542
	nrIndex           ID = 543
	nsoIndex          ID = 544
	nusIndex          ID = 545
	nusSSIndex        ID = 546
	nyIndex           ID = 547
	nynIndex          ID = 548
	nynUGIndex        ID = 549
	omIndex           ID = 550
	omETIndex         ID = 551
	omKEIndex         ID = 552
	orIndex           ID = 553
	orINIndex         ID = 554
	osIndex           ID = 555
	osGEIndex         ID = 556
	osRUIndex         ID = 557
	paIndex           ID = 558
	paArabIndex       ID = 559
	paArabPKIndex     ID = 560
	paGuruIndex       ID = 561
	paGuruINIndex     ID = 562
	papIndex          ID = 563
	plIndex           ID = 564
	plPLIndex         ID = 565
	prgIndex          ID = 566
	prg001Index       ID = 567
	psIndex           ID = 568
	psAFIndex         ID = 569
	ptIndex           ID = 570
	ptAOIndex         ID = 571
	ptBRIndex         ID = 572
	ptCHIndex         ID = 573
	ptCVIndex         ID = 574
	ptGQIndex         ID = 575
	ptGWIndex         ID = 576
	ptLUIndex         ID = 577
	ptMOIndex         ID = 578
	ptMZIndex         ID = 579
	ptPTIndex         ID = 580
	ptSTIndex         ID = 581
	ptTLIndex         ID = 582
	quIndex           ID = 583
	quBOIndex         ID = 584
	quECIndex         ID = 585
	quPEIndex         ID = 586
	rmIndex           ID = 587
	rmCHIndex         ID = 588
	rnIndex           ID = 589
	rnBIIndex         ID = 590
	roIndex           ID = 591
	roMDIndex         ID = 592
	roROIndex         ID = 593
	rofIndex          ID = 594
	rofTZIndex        ID = 595
	ruIndex           ID = 596
	ruBYIndex         ID = 597
	ruKGIndex         ID = 598
	ruKZIndex         ID = 599
	ruMDIndex         ID = 600
	ruRUIndex         ID = 601
	ruUAIndex         ID = 602
	rwIndex           ID = 603
	rwRWIndex         ID = 604
	rwkIndex          ID = 605
	rwkTZIndex        ID = 606
	sahIndex          ID = 607
	sahRUIndex        ID = 608
	saqIndex          ID = 609
	saqKEIndex        ID = 610
	sbpIndex          ID = 611
	sbpTZIndex        ID = 612
	sdIndex           ID = 613
	sdPKIndex         ID = 614
	sdhIndex          ID = 615
	seIndex           ID = 616
	seFIIndex         ID = 617
	seNOIndex         ID = 618
	seSEIndex         ID = 619
	sehIndex          ID = 620
	sehMZIndex        ID = 621
	sesIndex          ID = 622
	sesMLIndex        ID = 623
	sgIndex           ID = 624
	sgCFIndex         ID = 625
	shIndex           ID = 626
	shiIndex          ID = 627
	shiLatnIndex      ID = 628
	shiLatnMAIndex    ID = 629
	shiTfngIndex      ID = 630
	shiTfngMAIndex    ID = 631
	siIndex           ID = 632
	siLKIndex         ID = 633
	skIndex           ID = 634
	skSKIndex         ID = 635
	slIndex           ID = 636
	slSIIndex         ID = 637
	smaIndex          ID = 638
	smiIndex          ID = 639
	smjIndex          ID = 640
	smnIndex          ID = 641
	smnFIIndex        ID = 642
	smsIndex          ID = 643
	snIndex           ID = 644
	snZWIndex         ID = 645
	soIndex           ID = 646
	soDJIndex         ID = 647
	soETIndex         ID = 648
	soKEIndex         ID = 649
	soSOIndex         ID = 650
	sqIndex           ID = 651
	sqALIndex         ID = 652
	sqMKIndex         ID = 653
	sqXKIndex         ID = 654
	srIndex           ID = 655
	srCyrlIndex       ID = 656
	srCyrlBAIndex     ID = 657
	srCyrlMEIndex     ID = 658
	srCyrlRSIndex     ID = 659
	srCyrlXKIndex     ID = 660
	srLatnIndex       ID = 661
	srLatnBAIndex     ID = 662
	srLatnMEIndex     ID = 663
	srLatnRSIndex     ID = 664
	srLatnXKIndex     ID = 665
	ssIndex           ID = 666
	ssyIndex          ID = 667
	stIndex           ID = 668
	svIndex           ID = 669
	svAXIndex         ID = 670
	svFIIndex         ID = 671
	svSEIndex         ID = 672
	swIndex           ID = 673
	swCDIndex         ID = 674
	swKEIndex         ID = 675
	swTZIndex         ID = 676
	swUGIndex         ID = 677
	syrIndex          ID = 678
	taIndex           ID = 679
	taINIndex         ID = 680
	taLKIndex         ID = 681
	taMYIndex         ID = 682
	taSGIndex         ID = 683
	teIndex           ID = 684
	teINIndex         ID = 685
	teoIndex          ID = 686
	teoKEIndex        ID = 687
	teoUGIndex        ID = 688
	tgIndex           ID = 689
	tgTJIndex         ID = 690
	thIndex           ID = 691
	thTHIndex         ID = 692
	tiIndex           ID = 693
	tiERIndex         ID = 694
	tiETIndex         ID = 695
	tigIndex          ID = 696
	tkIndex           ID = 697
	tkTMIndex         ID = 698
	tlIndex           ID = 699
	tnIndex           ID = 700
	toIndex           ID = 701
	toTOIndex         ID = 702
	trIndex           ID = 703
	trCYIndex         ID = 704
	trTRIndex         ID = 705
	tsIndex           ID = 706
	ttIndex           ID = 707
	ttRUIndex         ID = 708
	twqIndex          ID = 709
	twqNEIndex        ID = 710
	tzmIndex          ID = 711
	tzmMAIndex        ID = 712
	ugIndex           ID = 713
	ugCNIndex         ID = 714
	ukIndex           ID = 715
	ukUAIndex         ID = 716
	urIndex           ID = 717
	urINIndex         ID = 718
	urPKIndex         ID = 719
	uzIndex           ID = 720
	uzArabIndex       ID = 721
	uzArabAFIndex     ID = 722
	uzCyrlIndex       ID = 723
	uzCyrlUZIndex     ID = 724
	uzLatnIndex       ID = 725
	uzLatnUZIndex     ID = 726
	vaiIndex          ID = 727
	vaiLatnIndex      ID = 728
	vaiLatnLRIndex    ID = 729
	vaiVaiiIndex      ID = 730
	vaiVaiiLRIndex    ID = 731
	veIndex           ID = 732
	viIndex           ID = 733
	viVNIndex         ID = 734
	voIndex           ID = 735
	vo001Index        ID = 736
	vunIndex          ID = 737
	vunTZIndex        ID = 738
	waIndex           ID = 739
	waeIndex          ID = 740
	waeCHIndex        ID = 741
	woIndex           ID = 742
	woSNIndex         ID = 743
	xhIndex           ID = 744
	xogIndex          ID = 745
	xogUGIndex        ID = 746
	yavIndex          ID = 747
	yavCMIndex        ID = 748
	yiIndex           ID = 749
	yi001Index        ID = 750
	yoIndex           ID = 751
	yoBJIndex         ID = 752
	yoNGIndex         ID = 753
	yueIndex          ID = 754
	yueHansIndex      ID = 755
	yueHansCNIndex    ID = 756
	yueHantIndex      ID = 757
	yueHantHKIndex    ID = 758
	zghIndex          ID = 759
	zghMAIndex        ID = 760
	zhIndex           ID = 761
	zhHansIndex       ID = 762
	zhHansCNIndex     ID = 763
	zhHansHKIndex     ID = 764
	zhHansMOIndex     ID = 765
	zhHansSGIndex     ID = 766
	zhHantIndex       ID = 767
	zhHantHKIndex     ID = 768
	zhHantMOIndex     ID = 769
	zhHantTWIndex     ID = 770
	zuIndex           ID = 771
	zuZAIndex         ID = 772
	caESvalenciaIndex ID = 773
	enUSuvaposixIndex ID = 774
)

var coreTags = []language.CompactCoreInfo{ // 773 elements
	// Entry 0 - 1F
	0x00000000, 0x01600000, 0x016000d3, 0x01600162,
	0x01c00000, 0x01c00052, 0x02100000, 0x02100081,
	0x02700000, 0x02700070, 0x03a00000, 0x03a00001,
	0x03a00023, 0x03a00039, 0x03a00063, 0x03a00068,
	0x03a0006c, 0x03a0006d, 0x03a0006e, 0x03a00098,
	0x03a0009c, 0x03a000a2, 0x03a000a9, 0x03a000ad,
	0x03a000b1, 0x03a000ba, 0x03a000bb, 0x03a000ca,
	0x03a000e2, 0x03a000ee, 0x03a000f4, 0x03a00109,
	// Entry 20 - 3F
	0x03a0010c, 0x03a00116, 0x03a00118, 0x03a0011d,
	0x03a00121, 0x03a00129, 0x03a0015f, 0x04000000,
	0x04300000, 0x0430009a, 0x04400000, 0x04400130,
	0x04800000, 0x0480006f, 0x05800000, 0x05820000,
	0x05820032, 0x0585b000, 0x0585b032, 0x05e00000,
	0x05e00052, 0x07100000, 0x07100047, 0x07500000,
	0x07500163, 0x07900000, 0x07900130, 0x07e00000,
	0x07e00038, 0x08200000, 0x0a000000, 0x0a0000c4,
	// Entry 40 - 5F
	0x0a500000, 0x0a500035, 0x0a50009a, 0x0a900000,
	0x0a900053, 0x0a90009a, 0x0b200000, 0x0b200079,
	0x0b500000, 0x0b50009a, 0x0b700000, 0x0b720000,
	0x0b720033, 0x0b75b000, 0x0b75b033, 0x0d700000,
	0x0d700022, 0x0d70006f, 0x0d700079, 0x0d70009f,
	0x0db00000, 0x0db00035, 0x0db0009a, 0x0dc00000,
	0x0dc00107, 0x0df00000, 0x0df00132, 0x0e500000,
	0x0e500136, 0x0e900000, 0x0e90009c, 0x0e90009d,
	// Entry 60 - 7F
	0x0fa00000, 0x0fa0005f, 0x0fe00000, 0x0fe00107,
	0x10000000, 0x1000007c, 0x10100000, 0x10100064,
	0x10100083, 0x10800000, 0x108000a5, 0x10d00000,
	0x10d0002e, 0x10d00036, 0x10d0004e, 0x10d00061,
	0x10d0009f, 0x10d000b3, 0x10d000b8, 0x11700000,
	0x117000d5, 0x11f00000, 0x11f00061, 0x12400000,
	0x12400052, 0x12800000, 0x12b00000, 0x12b00115,
	0x12d00000, 0x12d00043, 0x12f00000, 0x12f000a5,
	// Entry 80 - 9F
	0x13000000, 0x13000081, 0x13000123, 0x13600000,
	0x1360005e, 0x13600088, 0x13900000, 0x13900001,
	0x1390001a, 0x13900025, 0x13900026, 0x1390002d,
	0x1390002e, 0x1390002f, 0x13900034, 0x13900036,
	0x1390003a, 0x1390003d, 0x13900042, 0x13900046,
	0x13900048, 0x13900049, 0x1390004a, 0x1390004e,
	0x13900050, 0x13900052, 0x1390005d, 0x1390005e,
	0x13900061, 0x13900062, 0x13900064, 0x13900065,
	// Entry A0 - BF
	0x1390006e, 0x13900073, 0x13900074, 0x13900075,
	0x13900076, 0x1390007c, 0x1390007d, 0x13900080,
	0x13900081, 0x13900082, 0x13900084, 0x1390008b,
	0x1390008d, 0x1390008e, 0x13900097, 0x13900098,
	0x13900099, 0x1390009a, 0x1390009b, 0x139000a0,
	0x139000a1, 0x139000a5, 0x139000a8, 0x139000aa,
	0x139000ae, 0x139000b2, 0x139000b5, 0x139000b6,
	0x139000c0, 0x139000c1, 0x139000c7, 0x139000c8,
	// Entry C0 - DF
	0x139000cb, 0x139000cc, 0x139000cd, 0x139000cf,
	0x139000d1, 0x139000d3, 0x139000d6, 0x139000d7,
	0x139000da, 0x139000de, 0x139000e0, 0x139000e1,
	0x139000e7, 0x139000e8, 0x139000e9, 0x139000ec,
	0x139000ed, 0x139000f1, 0x13900108, 0x1390010a,
	0x1390010b, 0x1390010c, 0x1390010d, 0x1390010e,
	0x1390010f, 0x13900110, 0x13900113, 0x13900118,
	0x1390011c, 0x1390011e, 0x13900120, 0x13900126,
	// Entry E0 - FF
	0x1390012a, 0x1390012d, 0x1390012e, 0x13900130,
	0x13900132, 0x13900134, 0x13900136, 0x1390013a,
	0x1390013d, 0x1390013e, 0x13900140, 0x13900143,
	0x13900162, 0x13900163, 0x13900165, 0x13c00000,
	0x13c00001, 0x13e00000, 0x13e0001f, 0x13e0002c,
	0x13e0003f, 0x13e00041, 0x13e00048, 0x13e00051,
	0x13e00054, 0x13e00057, 0x13e0005a, 0x13e00066,
	0x13e00069, 0x13e0006a, 0x13e0006f, 0x13e00087,
	// Entry 100 - 11F
	0x13e0008a, 0x13e00090, 0x13e00095, 0x13e000d0,
	0x13e000d9, 0x13e000e3, 0x13e000e5, 0x13e000e8,
	0x13e000ed, 0x13e000f2, 0x13e0011b, 0x13e00136,
	0x13e00137, 0x13e0013c, 0x14000000, 0x1400006b,
	0x14500000, 0x1450006f, 0x14600000, 0x14600052,
	0x14800000, 0x14800024, 0x1480009d, 0x14e00000,
	0x14e00052, 0x14e00085, 0x14e000ca, 0x14e00115,
	0x15100000, 0x15100073, 0x15300000, 0x153000e8,
	// Entry 120 - 13F
	0x15800000, 0x15800064, 0x15800077, 0x15e00000,
	0x15e00036, 0x15e00037, 0x15e0003a, 0x15e0003b,
	0x15e0003c, 0x15e00049, 0x15e0004b, 0x15e0004c,
	0x15e0004d, 0x15e0004e, 0x15e0004f, 0x15e00052,
	0x15e00063, 0x15e00068, 0x15e00079, 0x15e0007b,
	0x15e0007f, 0x15e00085, 0x15e00086, 0x15e00087,
	0x15e00092, 0x15e000a9, 0x15e000b8, 0x15e000bb,
	0x15e000bc, 0x15e000bf, 0x15e000c0, 0x15e000c4,
	// Entry 140 - 15F
	0x15e000c9, 0x15e000ca, 0x15e000cd, 0x15e000d4,
	0x15e000d5, 0x15e000e6, 0x15e000eb, 0x15e00103,
	0x15e00108, 0x15e0010b, 0x15e00115, 0x15e0011d,
	0x15e00121, 0x15e00123, 0x15e00129, 0x15e00140,
	0x15e00141, 0x15e00160, 0x16900000, 0x1690009f,
	0x16d00000, 0x16d000da, 0x16e00000, 0x16e00097,
	0x17e00000, 0x17e0007c, 0x19000000, 0x1900006f,
	0x1a300000, 0x1a30004e, 0x1a300079, 0x1a3000b3,
	// Entry 160 - 17F
	0x1a400000, 0x1a40009a, 0x1a900000, 0x1ab00000,
	0x1ab000a5, 0x1ac00000, 0x1ac00099, 0x1b400000,
	0x1b400081, 0x1b4000d5, 0x1b4000d7, 0x1b800000,
	0x1b800136, 0x1bc00000, 0x1bc00098, 0x1be00000,
	0x1be0009a, 0x1d100000, 0x1d100033, 0x1d100091,
	0x1d200000, 0x1d200061, 0x1d500000, 0x1d500093,
	0x1d700000, 0x1d700028, 0x1e100000, 0x1e100096,
	0x1e700000, 0x1e7000d7, 0x1ea00000, 0x1ea00053,
	// Entry 180 - 19F
	0x1f300000, 0x1f500000, 0x1f800000, 0x1f80009e,
	0x1f900000, 0x1f90004e, 0x1f90009f, 0x1f900114,
	0x1f900139, 0x1fa00000, 0x1fb00000, 0x20000000,
	0x200000a3, 0x20300000, 0x20700000, 0x20700052,
	0x20800000, 0x20a00000, 0x20a00130, 0x20e00000,
	0x20f00000, 0x21000000, 0x2100007e, 0x21200000,
	0x21200068, 0x21600000, 0x21700000, 0x217000a5,
	0x21f00000, 0x22300000, 0x22300130, 0x22700000,
	// Entry 1A0 - 1BF
	0x2270005b, 0x23400000, 0x234000c4, 0x23900000,
	0x239000a5, 0x24200000, 0x242000af, 0x24400000,
	0x24400052, 0x24500000, 0x24500083, 0x24600000,
	0x246000a5, 0x24a00000, 0x24a000a7, 0x25100000,
	0x2510009a, 0x25400000, 0x254000ab, 0x254000ac,
	0x25600000, 0x2560009a, 0x26a00000, 0x26a0009a,
	0x26b00000, 0x26b00130, 0x26d00000, 0x26d00052,
	0x26e00000, 0x26e00061, 0x27400000, 0x28100000,
	// Entry 1C0 - 1DF
	0x2810007c, 0x28a00000, 0x28a000a6, 0x29100000,
	0x29100130, 0x29500000, 0x295000b8, 0x2a300000,
	0x2a300132, 0x2af00000, 0x2af00136, 0x2b500000,
	0x2b50002a, 0x2b50004b, 0x2b50004c, 0x2b50004d,
	0x2b800000, 0x2b8000b0, 0x2bf00000, 0x2bf0009c,
	0x2bf0009d, 0x2c000000, 0x2c0000b7, 0x2c200000,
	0x2c20004b, 0x2c400000, 0x2c4000a5, 0x2c500000,
	0x2c5000a5, 0x2c700000, 0x2c7000b9, 0x2d100000,
	// Entry 1E0 - 1FF
	0x2d1000a5, 0x2d100130, 0x2e900000, 0x2e9000a5,
	0x2ed00000, 0x2ed000cd, 0x2f100000, 0x2f1000c0,
	0x2f200000, 0x2f2000d2, 0x2f400000, 0x2f400052,
	0x2ff00000, 0x2ff000c3, 0x30400000, 0x3040009a,
	0x30b00000, 0x30b000c6, 0x31000000, 0x31b00000,
	0x31b0009a, 0x31f00000, 0x31f0003e, 0x31f000d1,
	0x31f0010e, 0x32000000, 0x320000cc, 0x32500000,
	0x32500052, 0x33100000, 0x331000c5, 0x33a00000,
	// Entry 200 - 21F
	0x33a0009d, 0x34100000, 0x34500000, 0x345000d3,
	0x34700000, 0x347000db, 0x34700111, 0x34e00000,
	0x34e00165, 0x35000000, 0x35000061, 0x350000da,
	0x35100000, 0x3510009a, 0x351000dc, 0x36700000,
	0x36700030, 0x36700036, 0x36700040, 0x3670005c,
	0x367000da, 0x36700117, 0x3670011c, 0x36800000,
	0x36800052, 0x36a00000, 0x36a000db, 0x36c00000,
	0x36c00052, 0x36f00000, 0x37500000, 0x37600000,
	// Entry 220 - 23F
	0x37a00000, 0x38000000, 0x38000118, 0x38700000,
	0x38900000, 0x38900132, 0x39000000, 0x39000070,
	0x390000a5, 0x39500000, 0x3950009a, 0x39800000,
	0x3980007e, 0x39800107, 0x39d00000, 0x39d05000,
	0x39d050e9, 0x39d36000, 0x39d3609a, 0x3a100000,
	0x3b300000, 0x3b3000ea, 0x3bd00000, 0x3bd00001,
	0x3be00000, 0x3be00024, 0x3c000000, 0x3c00002a,
	0x3c000041, 0x3c00004e, 0x3c00005b, 0x3c000087,
	// Entry 240 - 25F
	0x3c00008c, 0x3c0000b8, 0x3c0000c7, 0x3c0000d2,
	0x3c0000ef, 0x3c000119, 0x3c000127, 0x3c400000,
	0x3c40003f, 0x3c40006a, 0x3c4000e5, 0x3d400000,
	0x3d40004e, 0x3d900000, 0x3d90003a, 0x3dc00000,
	0x3dc000bd, 0x3dc00105, 0x3de00000, 0x3de00130,
	0x3e200000, 0x3e200047, 0x3e2000a6, 0x3e2000af,
	0x3e2000bd, 0x3e200107, 0x3e200131, 0x3e500000,
	0x3e500108, 0x3e600000, 0x3e600130, 0x3eb00000,
	// Entry 260 - 27F
	0x3eb00107, 0x3ec00000, 0x3ec000a5, 0x3f300000,
	0x3f300130, 0x3fa00000, 0x3fa000e9, 0x3fc00000,
	0x3fd00000, 0x3fd00073, 0x3fd000db, 0x3fd0010d,
	0x3ff00000, 0x3ff000d2, 0x40100000, 0x401000c4,
	0x40200000, 0x4020004c, 0x40700000, 0x40800000,
	0x4085b000, 0x4085b0bb, 0x408eb000, 0x408eb0bb,
	0x40c00000, 0x40c000b4, 0x41200000, 0x41200112,
	0x41600000, 0x41600110, 0x41c00000, 0x41d00000,
	// Entry 280 - 29F
	0x41e00000, 0x41f00000, 0x41f00073, 0x42200000,
	0x42300000, 0x42300165, 0x42900000, 0x42900063,
	0x42900070, 0x429000a5, 0x42900116, 0x43100000,
	0x43100027, 0x431000c3, 0x4310014e, 0x43200000,
	0x43220000, 0x43220033, 0x432200be, 0x43220106,
	0x4322014e, 0x4325b000, 0x4325b033, 0x4325b0be,
	0x4325b106, 0x4325b14e, 0x43700000, 0x43a00000,
	0x43b00000, 0x44400000, 0x44400031, 0x44400073,
	// Entry 2A0 - 2BF
	0x4440010d, 0x44500000, 0x4450004b, 0x445000a5,
	0x44500130, 0x44500132, 0x44e00000, 0x45000000,
	0x4500009a, 0x450000b4, 0x450000d1, 0x4500010e,
	0x46100000, 0x4610009a, 0x46400000, 0x464000a5,
	0x46400132, 0x46700000, 0x46700125, 0x46b00000,
	0x46b00124, 0x46f00000, 0x46f0006e, 0x46f00070,
	0x47100000, 0x47600000, 0x47600128, 0x47a00000,
	0x48000000, 0x48200000, 0x4820012a, 0x48a00000,
	// Entry 2C0 - 2DF
	0x48a0005e, 0x48a0012c, 0x48e00000, 0x49400000,
	0x49400107, 0x4a400000, 0x4a4000d5, 0x4a900000,
	0x4a9000bb, 0x4ac00000, 0x4ac00053, 0x4ae00000,
	0x4ae00131, 0x4b400000, 0x4b40009a, 0x4b4000e9,
	0x4bc00000, 0x4bc05000, 0x4bc05024, 0x4bc20000,
	0x4bc20138, 0x4bc5b000, 0x4bc5b138, 0x4be00000,
	0x4be5b000, 0x4be5b0b5, 0x4bef4000, 0x4bef40b5,
	0x4c000000, 0x4c300000, 0x4c30013f, 0x4c900000,
	// Entry 2E0 - 2FF
	0x4c900001, 0x4cc00000, 0x4cc00130, 0x4ce00000,
	0x4cf00000, 0x4cf0004e, 0x4e500000, 0x4e500115,
	0x4f200000, 0x4fb00000, 0x4fb00132, 0x50900000,
	0x50900052, 0x51200000, 0x51200001, 0x51800000,
	0x5180003b, 0x518000d7, 0x51f00000, 0x51f3b000,
	0x51f3b053, 0x51f3c000, 0x51f3c08e, 0x52800000,
	0x528000bb, 0x52900000, 0x5293b000, 0x5293b053,
	0x5293b08e, 0x5293b0c7, 0x5293b10e, 0x5293c000,
	// Entry 300 - 31F
	0x5293c08e, 0x5293c0c7, 0x5293c12f, 0x52f00000,
	0x52f00162,
} // Size: 3116 bytes

const specialTagsStr string = "ca-ES-valencia en-US-u-va-posix"

// Total table size 3147 bytes (3KiB); checksum: 5A8FFFA5
