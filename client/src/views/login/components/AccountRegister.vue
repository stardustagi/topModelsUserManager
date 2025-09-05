<template>
  <el-form
    ref="ruleFormRef"
    :model="ruleForm"
    :rules="updateRules"
    size="large"
  >
    <Motion>
      <el-form-item
        :rules="[
          {
            required: true,
            message: transformI18n($t('login.pureUsernameReg')),
            trigger: 'blur'
          }
        ]"
        prop="username"
      >
        <el-input
          v-model="ruleForm.username"
          clearable
          :placeholder="t('login.pureUsername')"
          :prefix-icon="useRenderIcon(User)"
        />
      </el-form-item>
    </Motion>

    <Motion :delay="200">
      <el-form-item prop="password">
        <el-input
          v-model="ruleForm.password"
          clearable
          show-password
          :placeholder="t('login.purePassword')"
          :prefix-icon="useRenderIcon(Lock)"
        />
      </el-form-item>
    </Motion>

    <Motion :delay="250">
      <el-form-item :rules="repeatPasswordRule" prop="repeatPassword">
        <el-input
          v-model="ruleForm.repeatPassword"
          clearable
          show-password
          :placeholder="t('login.pureSure')"
          :prefix-icon="useRenderIcon(Lock)"
        />
      </el-form-item>
    </Motion>

    <!--验证码-->
    <Motion :delay="200">
      <el-form-item prop="verifyCode">
        <el-input
          v-model="ruleForm.verifyCode"
          clearable
          :placeholder="t('login.pureVerifyCode')"
          :prefix-icon="useRenderIcon(Keyhole)"
        >
          <template v-slot:append>
            <ReImageVerify v-model:code="imgCode" />
          </template>
        </el-input>
      </el-form-item>
    </Motion>

    <!--
      <Motion :delay="300">
      <el-form-item>
        <el-checkbox v-model="checked">
          {{ t("login.pureReadAccept") }}
        </el-checkbox>
        <el-button link type="primary">
          {{ t("login.purePrivacyPolicy") }}
        </el-button>
      </el-form-item>
    </Motion>-->
    

    <Motion :delay="350">
      <el-form-item>
        <el-button
          class="w-full"
          size="default"
          type="primary"
          :loading="loading"
          @click="onUpdate(ruleFormRef)"
        >
          {{ t("login.pureDefinite") }}
        </el-button>
      </el-form-item>
    </Motion>

    <Motion :delay="400">
      <el-form-item>
        <el-button class="w-full" size="default" @click="onBack">
          {{ t("login.pureBack") }}
        </el-button>
      </el-form-item>
    </Motion>
  </el-form>
</template>


<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { ref, reactive } from "vue";
import Motion from "../utils/motion";
import { message } from "@/utils/message";
import { updateRules } from "../utils/rule";
import type { FormInstance } from "element-plus";
import { useVerifyCode } from "../utils/verifyCode";
import { $t, transformI18n } from "@/plugins/i18n";
import { useUserStoreHook } from "@/store/modules/user";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import Lock from "~icons/ri/lock-fill";
import User from "~icons/ri/user-3-fill";
import Keyhole from "~icons/ri/shield-keyhole-line";
import { ReImageVerify } from "@/components/ReImageVerify";
import { accountLoginApi } from "@/api/managerApi";

const { t } = useI18n();
const checked = ref(false);
const loading = ref(false);
const ruleForm = reactive({
  username: "",
  phone: "",
  verifyCode: "",
  password: "",
  repeatPassword: ""
});
const imgCode = ref("");
const ruleFormRef = ref<FormInstance>();
const { isDisabled, text } = useVerifyCode();
const repeatPasswordRule = [
  {
    validator: (rule, value, callback) => {
      if (value === "") {
        callback(new Error(transformI18n($t("login.purePassWordSureReg"))));
      } else if (ruleForm.password !== value) {
        callback(
          new Error(transformI18n($t("login.purePassWordDifferentReg")))
        );
      } else {
        callback();
      }
    },
    trigger: "blur"
  }
];

const onUpdate = async (formEl: FormInstance | undefined) => {
  loading.value = true;
  if (!formEl) return;
  await formEl.validate(valid => {
    if (valid) {
      // 暂时去掉勾选隐私政策
      if (true || checked.value) {
        const tempTime = useUserStoreHook()?.imageCodeTime || ""
        console.log("-------------- ", tempTime);
        accountLoginApi({ user_name: ruleForm.username, password: ruleForm.password, graph_verify_code: ruleForm.verifyCode, t: tempTime }).then((res) => {
          loading.value = false;
          console.log("登录数据 res ====================== ", res)
          if (res.errcode === 0) {
            message("注册成功", { type: "success" })
            setTimeout(() => {
              onBack()
            })
            console.log("login res ==== ", res);
          }
        }).catch((err) => {
          console.log(err);
          loading.value = false;
        })
        
        // 模拟请求，需根据实际开发进行修改
        // setTimeout(() => {
        //   message(transformI18n($t("login.pureRegisterSuccess")), {
        //     type: "success"
        //   });
        //   loading.value = false;
        // }, 2000);
      } else {
        loading.value = false;
        message(transformI18n($t("login.pureTickPrivacy")), {
          type: "warning"
        });
      }
    } else {
      loading.value = false;
    }
  });
};

function onBack() {
  useVerifyCode().end();
  useUserStoreHook().SET_CURRENTPAGE(0);
}
</script>

