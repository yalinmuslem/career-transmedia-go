<?php if (!defined('BASEPATH')) exit('No direct script access allowed'); ?>

<div class="container mb-15">
    <ol class="breadcrumb breadcrumb-dot fs-6 fw-semibold mb-10 text-muted">
        <li class="breadcrumb-item">
            <a href="<?= $this->login ? base_url(index_page() . '/userboard') : base_url(index_page()) ?>" class="">
                <span class="svg-icon svg-icon-3 svg-icon-muted svg-icon-primary"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M11 2.375L2 9.575V20.575C2 21.175 2.4 21.575 3 21.575H9C9.6 21.575 10 21.175 10 20.575V14.575C10 13.975 10.4 13.575 11 13.575H13C13.6 13.575 14 13.975 14 14.575V20.575C14 21.175 14.4 21.575 15 21.575H21C21.6 21.575 22 21.175 22 20.575V9.575L13 2.375C12.4 1.875 11.6 1.875 11 2.375Z" fill="currentColor" />
                    </svg>
                </span>
            </a>
        </li>
        <li class="breadcrumb-item fw-semibold text-primary"><?= $data['jabatan'] ?></li>
    </ol>

    <div class="row">
        <div class="d-flex justify-content-between">
            <div>
                <h1 class="fs-2hx fw-bolder mb-5"><?= $data['jabatan'] ?></h1>
                <div class="d-flex gap-2 mb-5">
                    <div class="align-items-center d-flex flex-row">
                        <div class="me-1">
                            <span class="svg-icon svg-icon-1x svg-icon-primary">
                                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path d="M8.7 4.19995L4 6.30005V18.8999L8.7 16.8V19L3.1 21.5C2.6 21.7 2 21.4 2 20.8V6C2 5.4 2.3 4.89995 2.9 4.69995L8.7 2.09998V4.19995Z" fill="currentColor" />
                                    <path d="M15.3 19.8L20 17.6999V5.09992L15.3 7.19989V4.99994L20.9 2.49994C21.4 2.29994 22 2.59989 22 3.19989V17.9999C22 18.5999 21.7 19.1 21.1 19.3L15.3 21.8998V19.8Z" fill="currentColor" />
                                    <path opacity="0.3" d="M15.3 7.19995L20 5.09998V17.7L15.3 19.8V7.19995Z" fill="currentColor" />
                                    <path opacity="0.3" d="M8.70001 4.19995V2L15.4 5V7.19995L8.70001 4.19995ZM8.70001 16.8V19L15.4 22V19.8L8.70001 16.8Z" fill="currentColor" />
                                    <path opacity="0.3" d="M8.7 16.8L4 18.8999V6.30005L8.7 4.19995V16.8Z" fill="currentColor" />
                                </svg>
                            </span>
                        </div>
                        <div class="text-primary text-uppercase"><?= $data['penempatan'] ?></div>
                    </div>
                    <div class="align-items-center d-flex flex-row">
                        <div class="me-1">
                            <span class="svg-icon svg-icon-1x svg-icon-primary">
                                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path d="M6.28548 15.0861C7.34369 13.1814 9.35142 12 11.5304 12H12.4696C14.6486 12 16.6563 13.1814 17.7145 15.0861L19.3493 18.0287C20.0899 19.3618 19.1259 21 17.601 21H6.39903C4.87406 21 3.91012 19.3618 4.65071 18.0287L6.28548 15.0861Z" fill="currentColor" />
                                    <rect opacity="0.3" x="8" y="3" width="8" height="8" rx="4" fill="currentColor" />
                                </svg>
                            </span>
                        </div>
                        <div class="text-primary text-uppercase"><?= $data['kelamin'] ?></div>
                    </div>
                    <div class="align-items-center d-flex flex-row">
                        <div class="me-1">
                            <span class="svg-icon svg-icon-1x svg-icon-primary">
                                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path opacity="0.3" d="M20 15H4C2.9 15 2 14.1 2 13V7C2 6.4 2.4 6 3 6H21C21.6 6 22 6.4 22 7V13C22 14.1 21.1 15 20 15ZM13 12H11C10.5 12 10 12.4 10 13V16C10 16.5 10.4 17 11 17H13C13.6 17 14 16.6 14 16V13C14 12.4 13.6 12 13 12Z" fill="currentColor" />
                                    <path d="M14 6V5H10V6H8V5C8 3.9 8.9 3 10 3H14C15.1 3 16 3.9 16 5V6H14ZM20 15H14V16C14 16.6 13.5 17 13 17H11C10.5 17 10 16.6 10 16V15H4C3.6 15 3.3 14.9 3 14.7V18C3 19.1 3.9 20 5 20H19C20.1 20 21 19.1 21 18V14.7C20.7 14.9 20.4 15 20 15Z" fill="currentColor" />
                                </svg>
                            </span>
                        </div>
                        <div class="text-primary text-uppercase"><?= $data['jurusan'] ?></div>
                    </div>
                </div>
            </div>
            <div>
                <a href="<?= $todo ?>" class="btn btn-pill btn-primary rounded-end-5 rounded-start-5">
                    <svg width="34" height="34" viewBox="0 0 34 34" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <g>
                            <path d="M17.0938 33.0156C8.15625 33.0156 0.796875 25.6562 0.796875 16.7188C0.796875 7.79688 8.15625 0.421875 17.0781 0.421875C26.0156 0.421875 33.3906 7.79688 33.3906 16.7188C33.3906 25.6562 26.0312 33.0156 17.0938 33.0156ZM25.5625 16.7188C25.5625 16.2031 25.375 15.8125 24.9688 15.4062L19.4062 10C19.0938 9.6875 18.7656 9.5625 18.3281 9.5625C17.5 9.5625 16.9062 10.1719 16.9062 11.0156C16.9062 11.4375 17.0938 11.8281 17.375 12.0938L19.2656 13.8906L21.0312 15.25L17.7344 15.0781H10.1719C9.25 15.0781 8.60938 15.7812 8.60938 16.7188C8.60938 17.6562 9.25 18.3438 10.1719 18.3438H17.7344L21.0312 18.1719L19.2969 19.5469L17.375 21.3281C17.0781 21.5938 16.9062 21.9844 16.9062 22.4062C16.9062 23.2656 17.5 23.875 18.3281 23.875C18.7656 23.875 19.0938 23.7344 19.4062 23.4375L24.9688 18.0312C25.3594 17.6406 25.5625 17.2344 25.5625 16.7188Z" fill="#FEFEFE" />
                        </g>
                    </svg>
                    Apply
                </a>
            </div>
        </div>
        <div class="col-md-8">
            <div class="card rounded-4">
                <div class="card-body">
                    <h1 class="fs-2 fw-bold mb-5">Description</h1>
                    <div class="mb-10">
                        <?= $data['Deskripsi'] ?>
                    </div>
                    <h1 class="fs-2 fw-bold mb-5">Job Responsibilities</h1>
                    <div>
                        <?= $data['Responsibility'] ?>
                    </div>
                </div>
            </div>
        </div>
        <div class="col-md-4">
            <div class="card rounded-4">
                <div class="card-body">
                    <div class="btn btn-primary rounded-end-5 rounded-start-5 mb-5">Other Position</div>
                    <?php foreach ($others as $value) : ?>
                        <div class="mb-3">
                            <div class="fs-4 mb-2 text-uppercase">
                                <a href="<?= base_url(index_page() . '/requirement-event/' . $value['uid']) ?>"><?= $value['jabatan'] ?></a>
                            </div>
                            <div><?= $value['penempatan'] ?>, <?= strtoupper($value['jurusan']) ?></div>
                            <div class="separator my-3"></div>
                        </div>
                    <?php endforeach ?>
                </div>
            </div>
        </div>
    </div>
</div>