<template>
    <div class="p-4">
        <el-card shadow="never">
            <el-table :data="list" border stripe v-loading="loading">
                <el-table-column prop="date" label="日期" width="120" />

                <el-table-column prop="provider_name" label="服务商" />

                <el-table-column prop="model_name" label="模型" />

                <el-table-column prop="consume_type" label="类型" width="100" />

                <el-table-column label="消费金额" align="right">
                    <template #default="{ row }">
                        {{ (row.total_cost / 10000).toFixed(4) }}
                    </template>
                </el-table-column>

                <!--
                <el-table-column label="成本 (万)" align="right">
                    <template #default="{ row }">
                        {{ (row.total_cost / 10000).toFixed(4) }}
                    </template>
                </el-table-column>
                -->
            </el-table>


            <el-pagination background layout="total, prev, pager, next" :total="total" :page-size="20"
                :current-page="currentPage" @current-change="onPageChange" />
        </el-card>
    </div>
</template>

<script setup lang="ts">
import { DayReportItem } from "@/api/apiParamsResp";
import { getDailyReportApi } from "@/api/managerApi";
import { ref, onMounted } from "vue";

const loading = ref(false);
const list = ref<DayReportItem[]>([]);
const total = ref(0);

const currentPage = ref(1);
const pageSize = ref(20);

const fetchData = async () => {
    loading.value = true;
    try {
        const skip = (currentPage.value - 1) * pageSize.value;
        const res = await getDailyReportApi({
            skip,
            limit: pageSize.value,
            sort: "id desc"
        });

        list.value = res.data.items || [];
        total.value = res.data.total || 0;
    } finally {
        loading.value = false;
    }
};

const onPageChange = (page: number) => {
    currentPage.value = page;
    fetchData();
};

const onSizeChange = (size: number) => {
    pageSize.value = size;
    currentPage.value = 1;
    fetchData();
};

onMounted(() => {
    fetchData();
});
</script>
